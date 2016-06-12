package main

import (
  "flag"
  "fmt"
  "log"
  "os"
  "os/exec"
  "strings"
  "syscall"
  "time"
)

func main() {
  var interval int
  flag.IntVar(&interval, "i", 1, "interval")

  var interrupt bool
  flag.BoolVar(&interrupt, "x", false, "interrupt")

  flag.Parse()

  command := flag.Args()

  loop(intervalToTime(interval), func () {
    reset()
    status := run(measure(command))
    if(interrupt && status != 0) {
      os.Exit(status)
    }
  })
}

func intervalToTime(i int) time.Duration {
  return time.Duration(i) * time.Second
}

func reset() {
  run([]string {"clear"})
}

func measure(command []string) []string {
  return append([]string {"time"}, command...)
}

func run(command []string) int {
  cmd := exec.Command(getShell(), wrapForShell(buildArgs(command))...)
  cmd.Env = os.Environ()
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Run()
  status, _ := cmd.ProcessState.Sys().(syscall.WaitStatus)
  return status.ExitStatus()
}

func safe(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func buildArgs(command []string) string {
  return strings.Join(command, " ")
}

func wrapForShell(command string) []string {
  return append([]string {"-c"}, fmt.Sprintf("%s; eval %s", sourceFiles(), command))
}

func sourceFiles() string {
  return fmt.Sprintf("%s %s", getShellSourceCmd(), getSourceFilePath())
}

func getSourceFilePath() string {
  return "~/.zshrc"
}

func getShellSourceCmd() string {
  return "."
}

func getShell() string {
  bin, found := syscall.Getenv("SHELL")
  if found == false {
    bin, _ = exec.LookPath("sh")
  }
  return bin
}

func loop(d time.Duration, fn func()) {
  fn()
  select {
  case <- time.After(d):
    loop(d, fn)
  }
}
