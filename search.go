package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
)

// searchToolOptions returns the structured email_search parameters, mirroring
// the structured search fields exposed by the mxmcp2 MCP server
// (internal/mcpsrv/server.go -> structuredSearchFields). The calling model fills
// these directly; the server skips its own LLM parse.
func searchToolOptions() []mcp.ToolOption {
	return []mcp.ToolOption{
		mcp.WithString("Query", mcp.Required(), mcp.Description(kQueryArgument)),
		mcp.WithString("SearchType",
			mcp.Enum("emails", "summary", "content", "stat", "seek"),
			mcp.Description(kSearchTypeArgument)),
		mcp.WithString("From", mcp.Description(kFromArgument)),
		mcp.WithString("To", mcp.Description(kToArgument)),
		mcp.WithString("StartDate", mcp.Description(kStartDateArgument)),
		mcp.WithString("EndDate", mcp.Description(kEndDateArgument)),
		mcp.WithString("Order",
			mcp.Enum("Newest", "Oldest"),
			mcp.Description(kOrderArgument)),
		mcp.WithString("BiDirection",
			mcp.Enum("Yes", "No"),
			mcp.Description(kBiDirectionArgument)),
		mcp.WithString("Words", mcp.Description(kWordsArgument)),
		mcp.WithString("K", mcp.Description(kKArgument)),
		mcp.WithString("MsgId", mcp.Description(kMsgIdArgument)),
		mcp.WithString("Thread",
			mcp.Enum("Yes", "No"),
			mcp.Description(kThreadArgument)),
		mcp.WithNumber("MaxTokens", mcp.Description(kMaxTokensArgument)),
		mcp.WithTitleAnnotation("Email Search"),
	}
}

// scope holds the structured search fields read from the tool request. The JSON
// tags match search.Scope on the mxmcp2 backend, which decodes this body in its
// /query/structured endpoint and runs QueryScope (no server-side LLM parse).
type scope struct {
	Query       string `json:"Query"`
	SearchType  string `json:"SearchType,omitempty"`
	From        string `json:"From,omitempty"`
	To          string `json:"To,omitempty"`
	StartDate   string `json:"StartDate,omitempty"`
	EndDate     string `json:"EndDate,omitempty"`
	Order       string `json:"Order,omitempty"`
	BiDirection string `json:"BiDirection,omitempty"`
	Words       string `json:"Words,omitempty"`
	K           string `json:"K,omitempty"`
	MsgId       string `json:"MsgId,omitempty"`
	Thread      string `json:"Thread,omitempty"`
	MaxTokens   int    `json:"MaxTokens,omitempty"`
}

func EmailSearchTool(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queryStr, queryErr := request.RequireString("Query")
	if queryErr != nil {
		return mcp.NewToolResultError("Query is required"), nil
	}

	s := scope{
		Query:       queryStr,
		SearchType:  request.GetString("SearchType", ""),
		From:        request.GetString("From", ""),
		To:          request.GetString("To", ""),
		StartDate:   request.GetString("StartDate", ""),
		EndDate:     request.GetString("EndDate", ""),
		Order:       request.GetString("Order", ""),
		BiDirection: request.GetString("BiDirection", ""),
		Words:       request.GetString("Words", ""),
		K:           request.GetString("K", ""),
		MsgId:       request.GetString("MsgId", ""),
		Thread:      request.GetString("Thread", ""),
		MaxTokens:   request.GetInt("MaxTokens", 0),
	}

	response, err := query(ctx, s)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return mcp.NewToolResultText(response), nil
}

func query(_ context.Context, s scope) (string, error) {
	client := &http.Client{}

	// Send the structured fields as a JSON body; the backend decodes them into a
	// search.Scope and runs the search directly (no server-side LLM parse).
	payload, err := json.Marshal(s)
	if err != nil {
		return "", fmt.Errorf("failed to encode request: %w", err)
	}

	endPointUrl := fmt.Sprintf("%s/query/structured", ApiHost)
	req, err := http.NewRequest(http.MethodPost, endPointUrl, bytes.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
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
