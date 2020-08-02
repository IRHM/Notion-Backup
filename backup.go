package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
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
	fmt.Println("Committing and pushing updated notes.")

	commitMsg := "Backup " + time.Now().Format("02.01.06 - 15:04:05") + ""

	// Commit and push all new updated notes in dir
	cmd := exec.Command("/bin/sh", "-c", "git add *;git commit -a -m '"+commitMsg+"'; git push")
	cmd.Dir = GitRepoFolder
	err := cmd.Run()
	if err != nil {
		fmt.Println("Couldn't commit backup: ", err)
		os.Exit(1)
	}

	fmt.Println("Pushed Commit '" + commitMsg + "'")
}
