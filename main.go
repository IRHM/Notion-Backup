package main

import "os"

func main() {
	// Download backup a get full path to it
	file := getBackup()

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
