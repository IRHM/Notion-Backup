package main

import (
	"os/exec"
	"runtime"
)

func runCommand(command string, dir string, stdin bool) (string, error) {
	shell := "/bin/sh"
	prefix := "-"

	// If on windows change shell to powershell
	if runtime.GOOS == "windows" {
		shell = "powershell"
		prefix = "/"
	}

	cmd := exec.Command(shell, prefix+"c", command)
	cmd.Dir = dir

	if stdin {
		cmd.StdinPipe()
	}

	out, err := cmd.CombinedOutput()

	return string(out), err
}
