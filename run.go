package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func reset() {
	run([]string{CLEAR_CMD})
}

func printStatus(s int) {
	fmt.Printf("\033[90mexit: %d\033[0m\n\n", s)
}

func makeCmd(command []string) *exec.Cmd {
	cmd := _makeCmd(command)
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
