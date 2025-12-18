package ssllabs

// AnalyzeResponse represents the main response from the /analyze API.
// It contains the high-level status of the scan.
type AnalyzeResponse struct {
	Host            string     `json:"host"`
	Port            int        `json:"port"`
	Protocol        string     `json:"protocol"`
	IsPublic        bool       `json:"isPublic"`
	Status          string     `json:"status"`
	StatusMessage   string     `json:"statusMessage,omitempty"`
	StartTime       int64      `json:"startTime"`
	TestTime        int64      `json:"testTime,omitempty"`
	EngineVersion   string     `json:"engineVersion"`
	CriteriaVersion string     `json:"criteriaVersion"`
	Endpoints       []Endpoint `json:"endpoints,omitempty"` // A domain can have multiple IP addresses
}

// Endpoint represents a specific server IP behind the domain.
type Endpoint struct {
	IPAddress         string `json:"ipAddress"`
	ServerName        string `json:"serverName,omitempty"`
	StatusMessage     string `json:"statusMessage,omitempty"`
	Grade             string `json:"grade,omitempty"` // e.g., "A+", "B", "F"
	GradeTrustIgnored string `json:"gradeTrustIgnored,omitempty"`
	HasWarnings       bool   `json:"hasWarnings"`
	IsExceptional     bool   `json:"isExceptional"`
	Progress          int    `json:"progress,omitempty"` // Progress percentage (0-100)
	Duration          int    `json:"duration,omitempty"`
	Delegation        int    `json:"delegation"`
}

// Constants for Status checks
const (
	StatusDns        = "DNS"
	StatusError      = "ERROR"
	StatusInProgress = "IN_PROGRESS"
	StatusReady      = "READY"
)
