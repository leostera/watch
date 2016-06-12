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
    info(command)
    run(command)
  })
}

func reset() {
  run([]string { "clear "})
}

func info(command []string) {
  fmt.Println("Running:", buildArgs(command), "\n")
}

func run(command []string) {
  cmd := exec.Command(getShell(), wrapForShell(buildArgs(command))...)
  cmd.Env = os.Environ()
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Run()
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
  return append([]string {"-c"}, fmt.Sprintf("eval %s", command))
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
