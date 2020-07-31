package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func getBackup() {
	enqueueTask()
}

func enqueueTask() {
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

	fmt.Println(t)

	req, err := http.NewRequest("POST", "https://www.notion.so/api/v3/enqueueTask", bytes.NewBuffer(reqBody))

	if err != nil {
		print("Error")
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:79.0) Gecko/20100101 Firefox/79.0")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")

	req.AddCookie(&http.Cookie{Name: "token_v2", Value: ""})

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		print("Error reading body. ")
	}

	fmt.Printf("%s\n", body)
}
