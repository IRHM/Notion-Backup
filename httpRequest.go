package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func post(url string, request []byte) string {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(request))

	if err != nil {
		return "error"
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "error"
	}

	return string(body)
}
