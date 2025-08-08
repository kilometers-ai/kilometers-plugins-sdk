package kmsdk

import "time"

// HTTP client types for API logging plugins

// APIClient provides HTTP client functionality for plugins
type APIClient interface {
	SendBatchEvent(events []BatchEventDto) error
}

// McpEventDto represents an MCP event for API logging
type McpEventDto struct {
	Timestamp     time.Time `json:"timestamp"`
	SessionID     string    `json:"session_id"`
	CorrelationID string    `json:"correlation_id"`
	EventType     string    `json:"event_type"`
	Direction     string    `json:"direction"`
	Data          string    `json:"data"` // Base64 encoded data
	Size          int       `json:"size"`
	CLIVersion    string    `json:"cli_version"`
}

// BatchEventDto represents a batch of events for API logging
type BatchEventDto struct {
	Events    []McpEventDto `json:"events"`
	Timestamp time.Time     `json:"timestamp"`
	BatchSize int           `json:"batch_size"`
}

// BatchRequest represents a batch request
type BatchRequest struct {
	Events []BatchEventDto `json:"events"`
}