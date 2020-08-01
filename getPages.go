package main

import (
	"encoding/json"
)

type sharedPagesRequest struct {
	IncludeDeleted bool `json:"includeDeleted"`
}

type sharedPagesResponse struct {
	Pages []struct {
		ID      string `json:"id"`
		SpaceID string `json:"spaceId"`
	} `json:"pages"`
}

func getPages() []string {
	t := sharedPagesRequest{
		IncludeDeleted: false,
	}

	// Serialize json request
	reqBody, err := json.Marshal(t)
	if err != nil {
		print("error")
	}

	// Send request
	reply := requestData("getUserSharedPages", reqBody)

	// Deserialize json response
	end := sharedPagesResponse{}
	json.Unmarshal(reply, &end)

	var pages []string
	for i := range end.Pages {
		// Add page block ids to pages[]
		pages = append(pages, end.Pages[i].ID)
	}

	return pages
}
