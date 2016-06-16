//+build windows

package main

import (
	"os/exec"
)

const CLEAR_CMD = "cls"

func _makeCmd(command []string) *exec.Cmd {
	return exec.Command("cmd", append([]string{"/C"}, command...)...)
}
