package kmsdk

import "context"

// KilometersPlugin defines the interface that all Kilometers plugins must implement
// This is the universal plugin interface for standalone plugins with JWT authentication
type KilometersPlugin interface {
	// Authenticate validates the plugin's credentials and permissions
	Authenticate(ctx context.Context, token string) error

	// ProcessMessage handles an MCP message and returns any events to log
	ProcessMessage(ctx context.Context, message []byte, direction string) ([]Event, error)

	// GetInfo returns plugin metadata
	GetInfo() PluginInfo
}