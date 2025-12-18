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
		fmt.Println("Usage: go run cmd/cli/main.go <domain> [new]")
		os.Exit(1)
	}

	// Get the domain and the flag for new scan, if new is present, set the flag to true
	domain := os.Args[1]
	newScan := ""
	if len(os.Args) > 2 {
		newScan = os.Args[2]
	}

	// Initialize the client
	client := ssllabs.NewClient()
	var resp *ssllabs.AnalyzeResponse
	var err error

	// Initiate the scan either new or cached
	switch newScan {
	case "new":
		fmt.Printf("🔍 Starting new scan for: %s\n", domain)
		resp, err = client.FreshAnalyze(domain)
	case "newsync":
		fmt.Printf("🔍 Starting new synchronous scan for: %s\n", domain)
		_, err = client.FreshAnalyze(domain)

		if err != nil {
			log.Fatalf("❌ Error starting fresh scan: %v", err)
		}

		// Wait for the scan to finish
		resp, err = client.WaitForResults(domain)
	default:
		fmt.Printf("🔍 Starting scan for: %s\n", domain)
		resp, err = client.Analyze(domain)
	}

	// Handle the error
	if err != nil {
		log.Fatalf("❌ Error analyzing domain: %v", err)
	}

	// Print the result nicely, using json.MarshalIndent to make it readable in the terminal
	prettyJSON, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(prettyJSON))
}
