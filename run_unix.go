//+build !windows

package main

import (
	"os/exec"
)

// ClearCmd is the command used to clean the terminal
const ClearCmd = "clear"

func _makeCmd(command []string) *exec.Cmd {
	path, _ := exec.LookPath(command[0])
	return exec.Command(path, command[1:]...)
}
