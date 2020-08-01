package main

import "fmt"

func main() {
	//getSettings()

	// Download backup a get full path to it
	files := getBackup()

	for _, val := range files {
		fmt.Println(val)
	}

	// // Extract downloaded backup
	// err := extract(file, "./export")
	// if err != nil {
	// 	print(err.Error)
	// }

	// // Delete zip
	// err = os.Remove(file)
	// if err != nil {
	// 	print(err.Error)
	// }
}
