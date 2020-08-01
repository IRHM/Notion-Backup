package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type settings struct {
	APITokenV2 string `json:"API_TOKEN_V2"`
}

// APITokenV2 setting
var APITokenV2 string

func getSettings() error {
	file := getSettingsFile()

	s, err := readSettings(file)
	if err != nil {
		return errors.New("Error getting settings: " + err.Error())
	}

	if len(s) < 1 {
		// If settings file is empty then ask user for settings and write them
		err := writeSettings(file)
		if err != nil {
			fmt.Println(err)
		}

		s, err = readSettings(file)
	}

	// Deserialize json settings
	end := settings{}
	json.Unmarshal(s, &end)

	APITokenV2 = end.APITokenV2

	return nil
}

func readSettings(file string) ([]byte, error) {
	// Read settings file
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return []byte(""), err
	}

	return data, nil
}

func writeSettings(file string) error {
	// Get user input for api token
	fmt.Println("API token_v2: ")
	var APITokenV2Input string
	fmt.Scanln(&APITokenV2Input)

	s := &settings{
		APITokenV2: APITokenV2Input,
	}

	// Serialize json request
	settingsJSON, err := json.Marshal(s)
	if err != nil {
		return err
	}

	// Write settings
	ioutil.WriteFile(file, settingsJSON, 0644)

	return nil
}

func getSettingsFile() string {
	var settingsFile = "settings.json"

	// Create settingsFile if it doesn't already exist
	if _, err := os.Stat(settingsFile); os.IsNotExist(err) {
		fmt.Println("Settings file not found, creating it.")
		os.Create(settingsFile)
	} else {
		fmt.Println("Found settings file.")
	}

	return settingsFile
}
