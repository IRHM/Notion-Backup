package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var nogit bool = false
	if os.Args[1] == "nogit" {
		nogit = true
		fmt.Println("Ok, won't backup to git. Will still move exported files to GitRepoFolder.")
	}

	// Get settings from settings.json
	// If they don't exist, create them
	err := getSettings()
	if err != nil {
		fmt.Println("Error getting settings: ", err)
	}

	if !nogit {
		// Check if git is accessible
		if !checkForGit() {
			fmt.Println("Can't access git, make sure it is in your $PATH.")
			os.Exit(1)
		}

		// Check if GitRepoFolder exists and is a git repo
		setupRepo()
	}

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

	if !nogit {
		// Commit and push changes
		commitBackup()
	}

	fmt.Println("Done backing up.")
}
