package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"io"
	"log"
	"net/http"
	"time"
)

func EmailSearchTool(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queryStr, ok := request.Params.Arguments["query"].(string)
	if !ok {
		return nil, errors.New("query must be a string")
	}

	response, err := query(ctx, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return mcp.NewToolResultText(response), nil
}

func query(_ context.Context, query string) (string, error) {
	client := &http.Client{}

	endPointUrl := fmt.Sprintf("%s/gpt/email/query", ApiHost)
	req, err := http.NewRequest("GET", endPointUrl, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Add query parameter
	q := req.URL.Query()
	q.Add("q", query)
	q.Add("t", time.Now().Local().Format(time.RFC3339))

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
