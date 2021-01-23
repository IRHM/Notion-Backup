package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var nogit bool = false

	// Check if any args have been set
	if len(os.Args) > 1 {
		if os.Args[1] == "nogit" {
			nogit = true
			fmt.Println("Ok, won't backup to git. Will still move exported files to GitRepoFolder.")
		} else {
			fmt.Println("The only supported arg is 'nogit'")
			os.Exit(1)
		}
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
			fmt.Println("Can't access git, make sure it is in your $PATH variable.")
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
