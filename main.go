package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	mcpserver "github.com/ThinkInAIXYZ/go-mcp/server"
	"github.com/ThinkInAIXYZ/go-mcp/transport"

	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	ApiHost  string
	UserName string
	Token    string
)

type UserData struct {
	Name string `json:"name" jsonschema:"required,description=The user's name'"`
}

func main() {
	var tool *protocol.Tool
	var toolErr error
	var err error

	log.SetFlags(log.Lshortfile)

	// Define flags
	defaultDescription := `Search query for emails from multiple email accounts.

Args:
- query: The user's query or statement related to email. Create a full sentence that best reflects what the user wants regarding emails. If the user is requesting a specific email, refer to tne notes below.

Example:
- query: "Get me the email from Bob about the new product"

Return:
It will return details about the search results plus 
a list of emails that match the query - if any.

Each email entry returned is a JSON object with the following fields:
- content: the email content
- subject: the email subject
- from: the email sender
- to: the email recipient
- date: the email date
- msgId: the email id
- link: a URL to view the email

Notes:
- When the user requests a specific email and you have the msgId,
indicate and use the msgId in your request or simply provide the link, if you have it.
- Do not assume the user is asking about their own email.`

	description := flag.String("d", defaultDescription, "Tool description")
	flag.StringVar(&Token, "t", "", "Token for authentication")

	flag.Parse()

	if Token == "" {
		// TODO REMOVE THIS IN FAVOR OF INIT RESPONSE
		panic("Token missing - required for authentication")
	}

	// Handle a graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	done := make(chan struct{})

	transportServer := transport.NewStdioServerTransport()

	mcpServer, serverErr := mcpserver.NewServer(transportServer)
	if serverErr != nil {
		log.Fatalf("Failed to create MCP server: %v", serverErr)
	}

	// Register the init tool
	//initDescription := `This must be called once before using any of the tools in this MCP.`
	//tool, toolErr = protocol.NewTool("init", initDescription, InitArgs{})
	//if toolErr != nil {
	//	log.Fatalf("Failed to create tool: %v", toolErr)
	//	return
	//}
	//mcpServer.RegisterTool(tool, mcpInit)

	// Register the search tool
	tool, toolErr = protocol.NewTool("email_search", *description, SimpleSearch{})
	if toolErr != nil {
		log.Fatalf("Failed to create tool: %v", toolErr)
		return
	}
	mcpServer.RegisterTool(tool, DevEmailSearch)

	// TODO port save user name to a separate tool

	go func() {
		<-c // Wait for a signal
		log.Println("Shutting down...")
		close(done) // Signal main to exit
		os.Exit(0)
	}()

	// Start server
	if err = mcpServer.Run(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

type InitArgs struct {
	User string `json:"user" jsonschema:"required,description=The user's name'"`
}

func mcpInit(req *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	// TODO USE THIS TO CHECK FOR A TOKEN
	var response string
	var initInstruction string // Placeholder
	var initReq InitArgs

	if err := protocol.VerifyAndUnmarshal(req.RawArguments, &initReq); err != nil {
		return nil, err
	}

	if UserName == "" {
		errMsg := "The user's name is required. DO NOT INFER THE USER NAME FROM PRIOR EMAILS. It must be provided by the user in chat and saved in the tool state. " + initInstruction
		return nil, fmt.Errorf(errMsg)
	} else {
		response = fmt.Sprintf("The user's name is: %s. %s", UserName, initInstruction)
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

// SaveUserDataToJSON saves the provided UserData to a file named "saved.json"
// in the current directory as JSON.
func SaveUserDataToJSON(userData UserData) error {
	// Marshal the UserData to JSON with indentation for better readability
	jsonData, err := json.MarshalIndent(userData, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling UserData to JSON: %w", err)
	}

	// Write the JSON data to "saved.json" in the current directory
	err = os.WriteFile("saved.json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file saved.json: %w", err)
	}

	return nil
}

// LoadUserDataFromJSON reads the UserData from the "saved.json" file
// and returns the UserData struct and any error encountered.
func LoadUserDataFromJSON() (UserData, error) {
	var userData UserData

	// Read the contents of the saved.json file
	fileData, err := os.ReadFile("saved.json")
	if err != nil {
		return userData, fmt.Errorf("error reading saved.json file: %w", err)
	}

	// Unmarshal the JSON data into the UserData struct
	err = json.Unmarshal(fileData, &userData)
	if err != nil {
		return userData, fmt.Errorf("error unmarshaling JSON to UserData: %w", err)
	}

	return userData, nil
}
