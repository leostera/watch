package main

import (
  "fmt"
  "log"
  "os"
  "os/exec"
  "time"
  "syscall"
)

func main() {
  command := os.Args[1:]
  loop(1 * time.Second, func () {
    run(command)
  })
}

func run(command []string) {
  cmd := exec.Command(getShell(), buildArgs(command)...)
  fmt.Println(cmd.Args)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Run()
}

func safe(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func buildArgs(command []string) []string {
  return append([]string {"-c"}, command...)
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
