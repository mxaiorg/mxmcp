package main

import (
	"fmt"
	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"io"
	"log"
	"net/http"
)

type SimpleSearch struct {
	Query string `json:"query" jsonschema:"required,description=The user query related to email"`
}

func DevEmailSearch(req *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	var searchReq SimpleSearch
	if argErr := protocol.VerifyAndUnmarshal(req.RawArguments, &searchReq); argErr != nil {
		return nil, argErr
	}

	response, err := query(searchReq.Query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &protocol.CallToolResult{
		Content: []protocol.Content{
			protocol.TextContent{
				Type: "text",
				Text: response,
			},
		},
	}, nil
}

func query(query string) (string, error) {
	client := &http.Client{}

	endPointUrl := fmt.Sprintf("%s/gpt/email/query", ApiHost)
	req, err := http.NewRequest("GET", endPointUrl, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Add query parameter
	q := req.URL.Query()
	q.Add("q", query)
	req.URL.RawQuery = q.Encode()

	// Add Bearer token to the Authorization header
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", Token))

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read and return the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}
