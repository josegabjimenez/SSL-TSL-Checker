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