package main

import (
	"flag"
	"fmt"
	"time"
)

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
		die(0, Version)
	}

	command := flag.Args()

	if usage || len(command) == 0 {
		help()
		die(0, "")
	}

	if !commandExists(command) {
		die(0, "Executable not found in PATH")
	}

	interval, err := time.ParseDuration(i)
	if err != nil {
		die(0, "Invalid interval: try 1s, 1ms, 2h45m2s")
	}

	loop(interval, func() {
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
