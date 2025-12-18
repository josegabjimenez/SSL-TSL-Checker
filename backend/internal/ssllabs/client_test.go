package ssllabs

import (
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := NewClient()

	// Check if the result is what we expect
	if client == nil {
		t.Error("Expected client to be non-nil")
	}

	// Check if the internal HTTP client was initialized
	if client.http == nil {
		t.Error("Expected internal HTTP client to be initialized")
	}

	// Check if the timeout was set correctly
	if client.http.Timeout != 10*time.Second {
		t.Errorf("Expected timeout to be 10s, got %v", client.http.Timeout)
	}
}
