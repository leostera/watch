package main

import (
  "fmt"
  "strings"
  "log"
  "os"
  "os/exec"
  "time"
)

func main() {
  command := buildCommand(os.Args[1:])
  loop(1 * time.Second, func () {
    run(command)
  })
}

func buildCommand(args []string) []string {
  cmd := fmt.Sprintf(`"%s"`, strings.Join(args, " "))
  return []string { "-c", cmd }
}

func run(command []string) {
  fmt.Println(command)
  cmd := exec.Command("/bin/sh", command...)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Run()
}

func safe(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func loop(d time.Duration, fn func()) {
  select {
  case <- time.After(d):
    fn()
    loop(d, fn)
  }
}
