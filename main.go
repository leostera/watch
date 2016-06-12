package main

import (
  "fmt"
  "log"
  "os"
  "os/exec"
  "time"
  "syscall"
  "strings"
)

func main() {
  command := os.Args[1:]
  loop(1 * time.Second, func () {
    reset()
    status := run(measure(command))
    if(status != 0) {
      os.Exit(status)
    }
  })
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
  select {
  case <- time.After(d):
    fn()
    loop(d, fn)
  }
}
