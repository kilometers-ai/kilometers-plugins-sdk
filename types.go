package kmsdk

import "time"

// Event represents a logged event from a plugin
type Event struct {
	ID        string                 `json:"id"`
	Timestamp string                 `json:"timestamp"`
	Type      string                 `json:"type"`
	Data      map[string]interface{} `json:"data"`
}

// PluginInfo contains metadata about a plugin
type PluginInfo struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	Description  string `json:"description"`
	RequiredTier string `json:"required_tier"`
}

// Config contains configuration for plugin initialization
type Config struct {
	ApiEndpoint string `json:"api_endpoint"`
	Debug       bool   `json:"debug"`
}

// Direction represents message direction
type Direction string

const (
	DirectionInbound  Direction = "inbound"
	DirectionOutbound Direction = "outbound"
)

// StreamEvent represents a stream event
type StreamEvent struct {
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}