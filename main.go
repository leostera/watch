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
		reset()
		status := run(command)
		if interrupt && status != 0 {
			die(status, "")
		}
	})
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

     -i, --interval             interval in seconds or ms, defaulting to 1s
     -x, --exit                 exit on failure
     -v, --version              print out version
     -h, --help                 this help page

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
	run([]string{"clear"})
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
