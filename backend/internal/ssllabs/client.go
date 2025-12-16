package ssllabs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Entry point for the SSL Labs API.
const BaseURL = "https://api.ssllabs.com/api/v2"

type Client struct {
	http *http.Client
}

// NewClient creates a usable instance of the SSL Labs client.
func NewClient() *Client {
	return &Client{
		http: &http.Client{
			Timeout: 10 * time.Second, // Default Go HTTP client has NO timeout.
		},
	}
}

// Analyze function triggers a TLS check for a specific domain.
// This will use cached results if available, or start a new assessment if needed.
func (c *Client) Analyze(domain string) (*AnalyzeResponse, error) {
	// 1. Construct the URL
	url := fmt.Sprintf("%s/analyze?host=%s&all=done", BaseURL, domain)

	// 2. Make the Request
	resp, err := c.http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to contact SSL Labs: %w", err)
	}
	
	// 3. Close body when this function finishes
	defer resp.Body.Close()

	// 4. Check HTTP Status Code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status: %d", resp.StatusCode)
	}

	// 5. Decode JSON directly into the Struct
	var result AnalyzeResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return &result, nil
}

// FreshAnalyze initiates a new SSL Labs assessment, ignoring cached results.
// This method should be called ONCE to start a new assessment. It could be called again when the assessment is complete.
func (c *Client) FreshAnalyze(domain string) (*AnalyzeResponse, error) {
	// 1. Construct the URL with startNew=on to force a new assessment
	url := fmt.Sprintf("%s/analyze?host=%s&startNew=on&all=done", BaseURL, domain)

	// 2. Make the Request
	resp, err := c.http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to contact SSL Labs: %w", err)
	}
	
	// 3. Close body when this function finishes
	defer resp.Body.Close()

	// 4. Check HTTP Status Code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status: %d", resp.StatusCode)
	}

	// 5. Decode JSON directly into the Struct
	var result AnalyzeResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return &result, nil
}

func (c *Client) WaitForResults(domain string) (*AnalyzeResponse, error) {
	// 1. Create a ticker that ticks every 5 seconds
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// 2. Set a timeout of 3 minutes
	timeout := time.After(3 * time.Minute)

	fmt.Println("⏳ Waiting for scan to complete...")

	// 3. Infinite Loop to check the status every 5 seconds
	for {
		select {
		case <-timeout:
			return nil, fmt.Errorf("timed out waiting for results")
		case <-ticker.C:
			// Check the status every 5 seconds
			resp, err := c.Analyze(domain)
			if err != nil {
				return nil, err
			}

			switch resp.Status {
			case StatusReady:
				return resp, nil
			case StatusError:
				return nil, fmt.Errorf("scan failed: %s", resp.StatusMessage)
			default:
				// Still working... print a dot to show that the scan is still running
				fmt.Print(".")
			}
		}
	}
}