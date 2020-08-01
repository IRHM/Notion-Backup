package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

func requestData(apiFile string, reqBody []byte) []byte {
	// Create new HTTP POST request
	req, err := http.NewRequest("POST", "https://www.notion.so/api/v3/"+apiFile, bytes.NewBuffer(reqBody))
	if err != nil {
		print("POST request error: ", err)
	}

	// Add HTTP headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:79.0) Gecko/20100101 Firefox/79.0")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")

	// Add api token in cookies
	req.AddCookie(&http.Cookie{Name: "token_v2", Value: APITokenV2})

	// Create client
	client := &http.Client{Timeout: time.Second * 10}

	// Execute request
	resp, err := client.Do(req)

	// Read response from req
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print("Error reading response body: ", err)
	}

	return body
}
