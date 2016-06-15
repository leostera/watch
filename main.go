package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const VERSION = "v0.0.1"

func main() {
	var i string
	flag.StringVar(&i, "i", "1s", "")
	flag.StringVar(&i, "interval", "1s", "")

	var interrupt bool
	flag.BoolVar(&interrupt, "x", false, "")
	flag.BoolVar(&interrupt, "exit", false, "")

	var clear bool
	flag.BoolVar(&clear, "c", false, "")
	flag.BoolVar(&clear, "clear", false, "")

	var version bool
	flag.BoolVar(&version, "v", false, "")
	flag.BoolVar(&version, "version", false, "")

	var usage bool
	flag.BoolVar(&usage, "h", false, "")
	flag.BoolVar(&usage, "help", false, "")

	flag.Parse()

	if version {
		die(0, VERSION)
	}

	command := flag.Args()

	if usage || len(command) == 0 {
		help()
		die(0, "")
	}

	if !commandExists(command) {
		die(0, "Executable not found in PATH")
	}

	interval := parseInterval(i)

	loop(intervalToTime(interval), func() {
		if clear {
			reset()
		}
		status := run(command)
		printStatus(status)
		if interrupt && status != 0 {
			die(status, "")
		}
	})
}

func printStatus(s int) {
	fmt.Printf("\033[90mexit: %d\033[0m\n\n", s)
}

func intervalSuffixes() []string {
	return []string{"ms", "s", "MS", "S"}
}

func unitToFactor(s string) float64 {
	if s == "ms" || s == "MS" {
		return 1.0
	}
	if s == "s" || s == "S" {
		return 1000.0
	}
	return 1000.0
}

func suffixToInterval(s string, i string) (float64, bool) {
	is_unit := strings.HasSuffix(i, s)
	if is_unit {
		n, err := strconv.ParseFloat(i[:len(i)-len(s)], 64)
		dieIf(err, -1, "Invalid interval")
		return n * unitToFactor(s), true
	}
	return 0, false
}

func parseInterval(i string) float64 {
	suffixes := intervalSuffixes()
	for j := 0; j < len(suffixes); j++ {
		if val, is := suffixToInterval(suffixes[j], i); is {
			return val
		}
	}
	return 0
}

func help() {
	s := `
   Usage: watch [options] <cmd>

   Sample: watch -i=100ms make

   Options:

     -c, --clear                clear screen between command runs
     -i, --interval             interval in seconds or ms, defaulting to 1s
     -x, --exit                 exit on failure

     -h, --help                 this help page
     -v, --version              print out version

`
	fmt.Print(s)
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
	run([]string{CLEAR_CMD})
}

func makeCmd(command []string) *exec.Cmd {
	cmd := makeCoreCmd(command)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func getExitStatus(err error) int {
	if err == nil {
		return 0
	}
	if exitErr, ok := err.(*exec.ExitError); ok {
		if s, ok := exitErr.ProcessState.Sys().(syscall.WaitStatus); ok {
			return s.ExitStatus()
		}
	}
	return -1
}

func commandExists(command []string) bool {
	_, result := exec.LookPath(command[0])
	return result == nil
}

func run(command []string) int {
	cmd := makeCmd(command)
	return getExitStatus(cmd.Run())
}

func loop(d time.Duration, fn func()) {
	fn()
	for range time.Tick(d) {
		fn()
	}
}
