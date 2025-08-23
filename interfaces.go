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

// Plugin defines the extended interface for SDK-style plugins
// This interface provides more granular control over plugin lifecycle and message handling
type Plugin interface {
	// GetInfo returns plugin metadata
	GetInfo() PluginInfo

	// Initialize initializes the plugin with configuration
	Initialize(ctx context.Context, config PluginConfig) error

	// HandleMessage handles a message with direction and correlation ID
	HandleMessage(ctx context.Context, data []byte, direction Direction, correlationID string) error

	// HandleError handles an error
	HandleError(ctx context.Context, err error)

	// HandleStreamEvent handles a stream event
	HandleStreamEvent(ctx context.Context, event StreamEvent)
}