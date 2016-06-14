package main

import (
  "flag"
  "fmt"
  "os"
  "os/exec"
  "strings"
  "strconv"
  "syscall"
  "time"
)

const VERSION="v0.0.1"

func main() {
  var i string
  flag.StringVar(&i, "i", "1s", "Interval to wait between executions in seconds")
  flag.StringVar(&i, "interval", "1s", "Interval to wait between executions in seconds")

  var interrupt bool
  flag.BoolVar(&interrupt, "x", false, "Exit and elevate status code if the command fails")
  flag.BoolVar(&interrupt, "exit", false, "Exit and elevate status code if the command fails")

  var version bool
  flag.BoolVar(&version, "v", false, "Print out version")
  flag.BoolVar(&version, "version", false, "Print out version")

  var usage bool
  flag.BoolVar(&usage, "h", false, "Print out version")
  flag.BoolVar(&usage, "help", false, "Print out version")

  flag.Parse()

  if version {
    die(0, VERSION)
  }

  command := flag.Args()

  if usage || len(command) == 0 {
    help()
    die(0, "")
  }

  if ! commandExists(command) {
    die(0, "Executable not found in PATH")
  }

  interval := parseInterval(i)

  loop(intervalToTime(interval), func () {
    reset()
    status := run(command)
    if(interrupt && status != 0) {
      die(status, "")
    }
  })
}

func intervalSuffixes() []string {
  return []string{ "ms", "s", "MS", "S" }
}

func unitToFactor(s string) float64 {
  if s == "ms" || s == "MS" { return 1.0 }
  if s == "S" || s == "S" { return 1000.0 }
  return 1000.0
}

func suffixToInterval (s string, i string) (float64, bool) {
  is_unit := strings.HasSuffix(i, s)
  if is_unit {
    n, err := strconv.ParseFloat(i[:len(i)-len(s)], 64)
    dieIf(err, -1, "Invalid interval")
    return n*unitToFactor(s), true
  }
  return 0, false
}

func parseInterval(i string) float64 {
  suffixes := intervalSuffixes()
  for j := 0; j < len(suffixes); j++ {
    if val, is := suffixToInterval(suffixes[j], i); is  {
      return val
    }
  }
  return 0
}

func help() {
  br := func () { fmt.Println("") }
  _p  := func (p string, t string) { fmt.Println(p, t) }
  p := func(t string) { _p("  ", t) }
  pp := func(t string) { _p("    ", t) }

  br()
  p("Usage: watch [options] <cmd>")
  br()
  p("Sample: watch -i=100ms make")
  br()
  p("Options:")
  br()
  pp("-i, --interval\t\tinterval in seconds or ms, defaulting to 1s")
  pp("-x, --exit\t\t\texit on failure")
  pp("-v, --version\t\tprint out version")
  pp("-h, --help\t\t\tthis help page")
  br()
}

func dieIf(err error, status int, message string) {
  if err != nil {
    die(status, message)
  }
}

func die(status int, message string) {
  if len(message) > 0 {
    fmt.Println(message)
  }
  os.Exit(status)
}

func intervalToTime(i float64) time.Duration {
  return time.Duration(i) * time.Millisecond
}

func reset() {
  run([]string {"clear"})
}

func makeCmd(bin string, args []string) *exec.Cmd {
  cmd := exec.Command(bin, args...)
  cmd.Env = os.Environ()
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  return cmd
}

func getExitStatus(cmd *exec.Cmd) int {
  status, _ := cmd.ProcessState.Sys().(syscall.WaitStatus)
  return status.ExitStatus()
}

func commandExists(command []string) bool {
  _, result := exec.LookPath(command[0])
  return result == nil
}

func run(command []string) int {
  path, _ := exec.LookPath(command[0])
  cmd := makeCmd(path, command[1:])
  cmd.Run()
  return getExitStatus(cmd)
}

func loop(d time.Duration, fn func()) {
  fn()
  for range time.Tick(d) {
    fn()
  }
}
