package main

import (
	"encoding/json"
	"fmt"
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

func getBackup() {
	taskID := enqueueTask()
	downloadExport(taskID)
}

// Returns taskId
func enqueueTask() string {
	t := &exportRequest{
		Task: &exportTask{
			EventName: "exportBlock",
			Request: &exportTaskRequest{
				BlockID:   "",
				Recursive: true,
				ExportOptions: &exportTaskRequestOptions{
					ExportType: "markdown",
					TimeZone:   "Europe/London",
					Locale:     "en",
				},
			},
		},
	}

	reqBody, err := json.Marshal(t)

	if err != nil {
		print("error")
	}

	reply := requestData("enqueueTask", reqBody)

	end := exportRequestResponse{}
	json.Unmarshal(reply, &end)

	return end.TaskID
}

func downloadExport(taskID string) {
	var exportURL string

	for {
		// Give some time for export file to be created
		time.Sleep(1000 * time.Millisecond)

		t := getTasks{
			TaskIDs: []string{taskID},
		}

		reqBody, err := json.Marshal(t)

		if err != nil {
			print("Error serializing json: ", err)
		}

		reply := requestData("getTasks", reqBody)

		end := getTasksResponse{}
		json.Unmarshal(reply, &end)

		if end.Results[0].Status.Type == "complete" {
			exportURL = end.Results[0].Status.ExportURL
			break
		}
	}

	fmt.Println(exportURL)
}
