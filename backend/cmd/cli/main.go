package main

import (
	"backend/internal/ssllabs"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// Get the domain from command line arguments
	// os.Args[0] is the program name, os.Args[1] is the first argument.
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/cli/main.go <domain>")
		os.Exit(1)
	}
	domain := os.Args[1]

	fmt.Printf("🔍 Starting scan for: %s\n", domain)

	// Initialize the client
	client := ssllabs.NewClient()

	// Call the Analyze function
	resp, err := client.Analyze(domain)
	if err != nil {
		log.Fatalf("❌ Error analyzing domain: %v", err)
	}

	// Print the result nicely, using json.MarshalIndent to make it readable in the terminal
	prettyJSON, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(prettyJSON))
}