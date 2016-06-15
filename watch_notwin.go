//+build !windows

package main

import (
	"os/exec"
)

const CLEAR_CMD = "clear"

func makeCoreCmd(command []string) *exec.Cmd {
	path, _ := exec.LookPath(command[0])
	return exec.Command(path, command[1:]...)
}
