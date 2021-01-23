package main

import (
	"fmt"
	"os"
	"strings"
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
	_, err = runCommand("git status", GitRepoFolder, false)
	if err != nil {
		fmt.Println("Folder is not initialized with git.")
		os.Exit(1)
	}
}

func checkForGit() bool {
	_, err := runCommand("git --version", "", false)
	if err != nil {
		return false
	}

	return true
}

func commitBackup() {
	fmt.Println("Committing and pushing updated notes.")

	commitMsg := "Backup " + time.Now().Format("02.01.06 - 15:04:05") + ""
	commands := [3]string{"git add *", "git commit -a -m \"" + commitMsg + "\"", "git push"}

	for _, command := range commands {
		out, err := runCommand(command, GitRepoFolder, true)
		if err != nil {
			outStr := strings.ToLower(string(out))

			if strings.Contains(outStr, "nothing to commit") {
				fmt.Println("No updated notes to commit.")
			} else if strings.Contains(outStr, "authentication failed") {
				fmt.Println("Authentication failed! Login and save credentials for git in the command line or follow authentication prompt.")
				os.Exit(1)
			} else {
				fmt.Println(fmt.Sprint(err) + ": " + string(out))
				os.Exit(1)
			}
		}
	}

	fmt.Println("Pushed Commit '" + commitMsg + "'")
}
