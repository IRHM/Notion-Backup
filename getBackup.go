package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

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
	data := []byte(`{}`)

	req, err := http.NewRequest("POST", "https://www.notion.so/api/v3/loadUserContent", bytes.NewBuffer(data))

	if err != nil {
		print("Error")
	}

	req.Header.Set("x-notion-active-user-header", "")

	req.AddCookie(&http.Cookie{Name: "__cfduid", Value: ""})
	req.AddCookie(&http.Cookie{Name: "notion_browser_id", Value: ""})
	req.AddCookie(&http.Cookie{Name: "notion_locale", Value: ""})
	req.AddCookie(&http.Cookie{Name: "intercom-id-gpfdrxfd", Value: ""})
	req.AddCookie(&http.Cookie{Name: "intercom-session-gpfdrxfd", Value: ""})
	req.AddCookie(&http.Cookie{Name: "token_v2", Value: ""})
	req.AddCookie(&http.Cookie{Name: "notion_user_id", Value: ""})
	req.AddCookie(&http.Cookie{Name: "notion_users", Value: ""})
	req.AddCookie(&http.Cookie{Name: "logglytrackingsession", Value: ""})

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print("Error reading body. ")
	}

	fmt.Printf("%s\n", body)

	//enqueueTask()
}

func enqueueTask() {
	t := task{
		eventName: "exportBlock",
		request: request{
			blockId:   "token",
			recursive: true,
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

	fmt.Printf("%+v\n", t)
	print("reqBody", reqBody)

	r := post(string("https://www.notion.so/api/v3/enqueueTask"), reqBody)

	print(r)
}
