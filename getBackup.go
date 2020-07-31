package main

import "encoding/json"

type task struct {
	eventName string
	request   request
}

type request struct {
	blockId       string
	recursive     bool
	exportOptions exportOptions
}

type exportOptions struct {
	exportType string
	timeZone   string
	locale     string
}

func getBackup() {
	enqueueTask()
}

func enqueueTask() {
	t := task{
		eventName: "exportBlock",
		request: request{
			blockId:   "token",
			recursive: false,
			exportOptions: exportOptions{
				exportType: "markdown",
				timeZone:   "Europe/London",
				locale:     "en",
			},
		},
	}

	reqBody, err := json.Marshal(t)

	if err != nil {
		print("error")
	}

	r := post(string("https://www.notion.so/api/v3/enqueueTask"), reqBody)

	print(r)
}
