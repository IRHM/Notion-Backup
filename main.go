package main

import (
	"fmt"
	"os"
)

func main() {
	err := getSettings()
	if err != nil {
		fmt.Println("Error getting settings: ", err)
	}

	// Download backup a get full path to it
	files := getBackup()

	// Foreach export zip, extract to 'export' folder then delete zip
	for _, file := range files {
		// Extract downloaded backup
		err := extract(file, "./export")
		if err != nil {
			print(err.Error)
		}

		// Delete zip
		err = os.Remove(file)
		if err != nil {
			print(err.Error)
		}
	}
}
