package handlers

import (
	"encoding/json"
	"backend/internal/ssllabs"
	"net/http"
)

// Handles requests to check a domain.
// Example: GET /api/scan?domain=google.com&new=true
func ScanHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Parse the query parameters (domain and new)
	domain := r.URL.Query().Get("domain")
	if domain == "" {
		http.Error(w, `{"error": "domain is required"}`, http.StatusBadRequest)
		return
	}
	
	shouldStartNew := r.URL.Query().Get("new") == "true"

	// Initialize the client 
	client := ssllabs.NewClient()
	var resp *ssllabs.AnalyzeResponse
	var err error

	// Initiate the scan either new or cached
	if shouldStartNew {
		resp, err = client.FreshAnalyze(domain)
	} else {
		resp, err = client.Analyze(domain)
	}

	// Handle the error
	if err != nil {
		http.Error(w, `{"error": "failed to scan domain"}`, http.StatusInternalServerError)
		return
	}

	// Send Response
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, `{"error": "failed to encode response"}`, http.StatusInternalServerError)
	}
}