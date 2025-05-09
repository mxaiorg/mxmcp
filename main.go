package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	// ApiHost is compiled in at build time
	ApiHost string
	// Token is currently provided by configuration
	Token    string
	UserName string
)

type UserData struct {
	Name string `json:"name" jsonschema:"required,description=The user's name'"`
}

func main() {
	log.SetFlags(log.Lshortfile)

	// Get flags
	description := flag.String("d", kSearchDescription, "Tool description")
	flag.StringVar(&Token, "t", "", "Token for authentication")
	flag.Parse()

	if Token == "" {
		// TODO REMOVE THIS IN FAVOR OF INIT RESPONSE
		panic("Token missing - required for authentication")
	}

	// Create a new MCP server
	s := server.NewMCPServer(
		"mxHERO MCP",
		"0.0.1",
		server.WithLogging(),
		server.WithRecovery(),
	)

	search := mcp.NewTool("email_search",
		mcp.WithDescription(*description),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("The user query related to email"),
		),
	)

	s.AddTool(search, EmailSearchTool)

	// Handle a graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	done := make(chan struct{})

	go func() {
		<-c // Wait for a signal
		log.Println("Shutting down...")
		close(done) // Signal main to exit
		os.Exit(0)
	}()

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
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
