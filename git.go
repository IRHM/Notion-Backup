package main

import (
	"fmt"
	"os"
	"os/exec"
)

func startBackup() {
	if !checkForGit() {
		fmt.Println("Can't access git, make sure it is in your $PATH.")
		os.Exit(1)
	}
}

func checkForGit() bool {
	cmd := exec.Command("git", "--version")
	err := cmd.Run()
	if err != nil {
		return false
	}

	return true
}
