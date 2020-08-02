package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type exportRequest struct {
	Task *exportTask `json:"task"`
}

type exportTask struct {
	EventName string             `json:"eventName"`
	Request   *exportTaskRequest `json:"request"`
}

type exportTaskRequest struct {
	BlockID       string                    `json:"blockId"`
	Recursive     bool                      `json:"recursive"`
	ExportOptions *exportTaskRequestOptions `json:"exportOptions"`
}

type exportTaskRequestOptions struct {
	ExportType string `json:"exportType"`
	TimeZone   string `json:"timeZone"`
	Locale     string `json:"locale"`
}

type exportRequestResponse struct {
	TaskID string `json:"taskId"`
}

type getTasks struct {
	TaskIDs []string `json:"taskIds"`
}

type getTasksResponse struct {
	Results []*getTasksResponseResults `json:"results"`
}

type getTasksResponseResults struct {
	Status *getTasksResponseStatus `json:"status"`
}

type getTasksResponseStatus struct {
	Type      string `json:"type"`
	ExportURL string `json:"exportURL"`
}

// Return all file paths of exported zips
func getBackup() []string {
	return enqueueTask()
}

func enqueueTask() []string {
	var exportedFiles []string
	blockIDS := getPages()

	for _, val := range blockIDS {
		fmt.Print("Working on: ", val)

		t := &exportRequest{
			Task: &exportTask{
				EventName: "exportBlock",
				Request: &exportTaskRequest{
					BlockID:   val,
					Recursive: true,
					ExportOptions: &exportTaskRequestOptions{
						ExportType: "markdown",
						TimeZone:   "Europe/London",
						Locale:     "en",
					},
				},
			},
		}

		// Serialize json request
		reqBody, err := json.Marshal(t)
		if err != nil {
			print("error")
		}

		// Send request
		reply := requestData("enqueueTask", reqBody)

		// Deserialize json response
		end := exportRequestResponse{}
		json.Unmarshal(reply, &end)

		// Download export and add its full path to 'exportedFiles'
		exportedFiles = append(exportedFiles, downloadExport(end.TaskID))

		fmt.Println(" ... Completed!")
	}

	return exportedFiles
}

func downloadExport(taskID string) string {
	var exportURL string

	// Get exportURL
	for {
		// Give some time for export file to be created
		time.Sleep(1000 * time.Millisecond)

		t := getTasks{
			TaskIDs: []string{taskID},
		}

		// Serialize json request
		reqBody, err := json.Marshal(t)
		if err != nil {
			print("Error serializing json: ", err)
		}

		// Request task info
		reply := requestData("getTasks", reqBody)

		// Deserialize json response
		end := getTasksResponse{}
		json.Unmarshal(reply, &end)

		// If status is complete, break from loop and set exportURL var
		if end.Results != nil {
			s := end.Results[0].Status
			if s != nil && s.Type == "complete" {
				exportURL = s.ExportURL
				break
			}
		}

		// Wait an extra second and tell user
		fmt.Println("Checking if file is done again in 2 seconds...")
		time.Sleep(1000 * time.Millisecond)
	}

	// Download file to 'export<randomstr>.zip'
	var filename string = "export" + randomStr(10) + ".zip"

	// Get the response bytes from the url
	resp, err := http.Get(exportURL)
	if err != nil {
		print("Error downloading file: ", err)
	}
	defer resp.Body.Close()

	// Create an empty file
	file, err := os.Create(filename)
	if err != nil {
		print("Error creating file: ", err)
	}
	defer file.Close()

	// Write the bytes to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		print("Error writing to file: ", err)
	}

	// Return downloaded files full path
	path, err := filepath.Abs(filename)
	if err != nil {
		print("Error getting exported zips absolute path ", err)
	}

	return path
}
