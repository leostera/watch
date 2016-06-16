//+build windows

package main

import (
	"os/exec"
)

// ClearCmd is the command used to clean the terminal
const ClearCmd = "cls"

func _makeCmd(command []string) *exec.Cmd {
	return exec.Command("cmd", append([]string{"/C"}, command...)...)
}
