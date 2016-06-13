package main

import (
  "flag"
  "fmt"
  "os"
  "os/exec"
  "strings"
  "syscall"
  "time"
)

const VERSION="v0.0.1"

func main() {
  var interval float64
  flag.Float64Var(&interval, "i", 1.0, "Interval to wait between executions in seconds")

  var interrupt bool
  flag.BoolVar(&interrupt, "x", false, "Exit and elevate status code if the command fails")

  var version bool
  flag.BoolVar(&version, "v", false, "Print out version")

  flag.Parse()

  if version {
    fmt.Println(VERSION)
    os.Exit(0)
  }

  command := flag.Args()

  if len(command) == 0 {
    os.Exit(0)
  }

  loop(intervalToTime(1000*interval), func () {
    status := run(measure(command))
    if(interrupt && status != 0) {
      os.Exit(status)
    }
  })
}

func intervalToTime(i float64) time.Duration {
  return time.Duration(i) * time.Millisecond
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

func buildArgs(command []string) string {
  return strings.Join(command, " ")
}

func wrapForShell(command string) []string {
  return []string {"-c", fmt.Sprintf("%s; eval %s", sourceFiles(), command)}
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
  for {
    select {
    case <- time.After(d):
      fn()
    }
  }
}
