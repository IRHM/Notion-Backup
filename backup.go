package main

import (
	"fmt"
	"os"
	"os/exec"
)

func setupRepo() {
	// Check if folder exists
	_, err := os.Stat(GitRepoFolder)
	if os.IsNotExist(err) {
		fmt.Println("Git repo folder doesn't exist. Edit your settings: " + getSettingsFile())
		os.Exit(1)
	}

	// Check if folder is a git repository
	cmd := exec.Command("git", "status")
	cmd.Dir = GitRepoFolder
	err = cmd.Run()
	if err != nil {
		fmt.Println("Folder is not initialized with git.")
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

func commitBackup() {
	cmd := exec.Command("git", "add *")
	cmd.Dir = GitRepoFolder
	err := cmd.Run()
	if err != nil {
		fmt.Println("Couldn't run 'git add *'")
		os.Exit(1)
	}
}
