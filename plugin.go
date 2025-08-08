package kmsdk

import "context"

// Plugin is the public contract all Kilometers plugins must implement.
type Plugin interface {
	// Authenticate validates credentials/permissions. The value is the host-provided token or API key.
	Authenticate(ctx context.Context, token string) error

	// Initialize provides host configuration and prepares the plugin for handling messages.
	Initialize(ctx context.Context, cfg Config) error

	// Shutdown releases resources.
	Shutdown(ctx context.Context) error

	// HandleMessage is called for each intercepted MCP message.
	HandleMessage(ctx context.Context, data []byte, direction Direction, correlationID string) error

	// HandleError informs the plugin of a non-fatal error in the host pipeline.
	HandleError(ctx context.Context, err error)

	// HandleStreamEvent informs the plugin of lifecycle events (start/stop/flush/etc).
	HandleStreamEvent(ctx context.Context, event StreamEvent)

	// GetInfo returns plugin metadata for discovery and display.
	GetInfo() PluginInfo
}
