package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Get settings from settings.json
	// If they don't exist, create them
	err := getSettings()
	if err != nil {
		fmt.Println("Error getting settings: ", err)
	}

	// Check if git is accessible
	if !checkForGit() {
		fmt.Println("Can't access git, make sure it is in your $PATH.")
		os.Exit(1)
	}

	// Check if GitRepoFolder exists and is a git repo
	setupRepo()

	// Download backup & get full path to it
	files := getBackup()

	// Foreach export zip, extract to gitrepo folder then delete zip
	for _, file := range files {
		// Extract downloaded backup
		err := extract(file, filepath.Join(GitRepoFolder, "notes"))
		if err != nil {
			print(err.Error)
		}

		// Delete zip
		err = os.Remove(file)
		if err != nil {
			print(err.Error)
		}
	}

	// Commit and push changes
	// commitBackup()
}
