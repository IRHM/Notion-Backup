package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type settings struct {
	APITokenV2    string `json:"API_TOKEN_V2"`
	GitRepoFolder string `json:"GIT_REPO_FOLDER"`
}

// APITokenV2 : Users token for notion api
var APITokenV2 string

// GitRepoFolder : Where should program find git repository
var GitRepoFolder string

func getSettings() error {
	// Get settings file, create if doesn't exist
	file := getSettingsFile()

	// Read everything in settings file
	s, err := readSettings(file)
	if err != nil {
		return errors.New("Error getting settings: " + err.Error())
	}

	// Deserialize json settings
	end := settings{}
	json.Unmarshal(s, &end)

	var didAppnd int = 0

	// Check if each setting is in json file, if they aren't ask for them and write them

	if end.APITokenV2 != "" {
		APITokenV2 = end.APITokenV2
	} else {
		APITokenV2Input := askForSetting("API Token: ")
		end.APITokenV2 = APITokenV2Input

		didAppnd++
	}

	if end.GitRepoFolder != "" {
		GitRepoFolder = end.GitRepoFolder
	} else {
		var GitRepoFolderInput string

		// Get git repo from user & keep asking until user gives a folder that exists
		for {
			GitRepoFolderInput = askForSetting("Git repo folder: ")

			_, err := os.Stat(GitRepoFolderInput)
			if os.IsNotExist(err) {
				fmt.Println("That folder does not exist. Make sure you are using the correct path!")
				GitRepoFolderInput = askForSetting("Git repo folder: ")
			} else {
				break
			}
		}

		end.GitRepoFolder = GitRepoFolderInput

		didAppnd++
	}

	// If did append settings, write all settings back to file
	if didAppnd > 0 {
		err := writeSettings(file, end)
		if err != nil {
			return err
		}

		// Re-run getSettings to set global setting vars
		getSettings()
	}

	return nil
}

func writeSettings(file string, s settings) error {
	// Serialize json request
	settingsJSON, err := json.Marshal(s)
	if err != nil {
		return errors.New("Error writing settings: " + err.Error())
	}

	// Write settings
	ioutil.WriteFile(file, settingsJSON, 0644)

	return nil
}

func askForSetting(question string) string {
	fmt.Println(question)
	var response string
	fmt.Scanln(&response)

	return response
}

func readSettings(file string) ([]byte, error) {
	// Read settings file
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return []byte(""), err
	}

	return data, nil
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
