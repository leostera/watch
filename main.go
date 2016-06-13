package main

import (
  "flag"
  "fmt"
  "os"
  "os/exec"
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
    reset()
    status := run(command)
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

func makeCmd(bin string, args []string) *exec.Cmd {
  path, _ := exec.LookPath(bin)
  cmd := exec.Command(path, args...)
  cmd.Env = os.Environ()
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  return cmd
}

func getExitStatus(cmd *exec.Cmd) int {
  status, _ := cmd.ProcessState.Sys().(syscall.WaitStatus)
  return status.ExitStatus()
}

func run(command []string) int {
  cmd := makeCmd(command[0], command[1:])
  cmd.Run()
  return getExitStatus(cmd)
}

func loop(d time.Duration, fn func()) {
  fn()
  for range time.Tick(d) {
    fn()
  }
}
