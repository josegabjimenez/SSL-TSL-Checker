package main

import (
	"backend/internal/handlers"
	"log"
	"net/http"
)

func main() {
	// Register Route /api/scan to handle the scan request
	http.HandleFunc("/api/scan", handlers.ScanHandler)

	// Start the server on port 8080
	port := ":8080"
	log.Printf("🚀 Server starting on http://localhost%s", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
