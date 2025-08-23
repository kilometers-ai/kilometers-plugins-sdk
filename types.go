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
	
	// Extended metadata for .kmpkg packages
	Author        string   `json:"author,omitempty"`
	CLIVersionMin string   `json:"cli_version_min,omitempty"`
	CLIVersionMax string   `json:"cli_version_max,omitempty"`
	Platforms     []string `json:"platforms,omitempty"`
}

// PluginConfig contains configuration for plugin initialization
// Consolidated from CLI's PluginConfig with enhanced ApiKey support
type PluginConfig struct {
	ApiEndpoint string `json:"api_endpoint"`
	Debug       bool   `json:"debug"`
	ApiKey      string `json:"api_key"`
}

// Config is an alias for PluginConfig for backward compatibility
type Config = PluginConfig

// Direction represents message direction
type Direction string

const (
	DirectionInbound  Direction = "inbound"
	DirectionOutbound Direction = "outbound"
)

// StreamEventType represents the type of stream event
// Consolidated from CLI's PluginStreamEventType with consistent naming
type StreamEventType string

const (
	StreamEventTypeStart StreamEventType = "start"
	StreamEventTypeEnd   StreamEventType = "end"
	StreamEventTypeError StreamEventType = "error"
)

// StreamEvent represents a stream event
// Enhanced structure consolidated from CLI's PluginStreamEvent
type StreamEvent struct {
	Type      StreamEventType   `json:"type"`
	Timestamp time.Time         `json:"timestamp"`
	Data      map[string]string `json:"data"`
}

// PluginError represents an error from plugin operations
// Consolidated from CLI for consistent error handling across plugins
type PluginError struct {
	Message string `json:"message"`
	Code    string `json:"code,omitempty"`
}

func (e *PluginError) Error() string {
	if e.Code != "" {
		return e.Code + ": " + e.Message
	}
	return e.Message
}

// NewPluginError creates a new plugin error
func NewPluginError(message string) *PluginError {
	return &PluginError{Message: message}
}

// NewPluginErrorWithCode creates a new plugin error with an error code
func NewPluginErrorWithCode(code, message string) *PluginError {
	return &PluginError{Code: code, Message: message}
}

// KmpkgMetadata represents the metadata contained within a .kmpkg package file
type KmpkgMetadata struct {
	// Core plugin information (extends PluginInfo)
	PluginInfo
	
	// Package-specific metadata
	BinaryName    string            `json:"binary_name"`
	Dependencies  []KmpkgDependency `json:"dependencies,omitempty"`
	Configuration map[string]string `json:"configuration,omitempty"`
	
	// Integrity and security
	Checksum  string    `json:"checksum"`
	Signature string    `json:"signature,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	
	// Installation metadata
	InstallInstructions string            `json:"install_instructions,omitempty"`
	Environment         map[string]string `json:"environment,omitempty"`
}

// KmpkgDependency represents a dependency required by a .kmpkg package
type KmpkgDependency struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"` // "plugin", "system", "library"
}

// KmpkgPackage represents a local .kmpkg plugin package
type KmpkgPackage struct {
	// File system information
	FilePath   string    `json:"file_path"`
	FileName   string    `json:"file_name"`
	FileSize   int64     `json:"file_size"`
	ModifiedAt time.Time `json:"modified_at"`
	
	// Package metadata (from manifest inside .kmpkg)
	Metadata KmpkgMetadata `json:"metadata"`
}

// IsCompatible checks if the package is compatible with the current environment
func (m *KmpkgMetadata) IsCompatible(cliVersion string, platform string) bool {
	// Check platform compatibility
	if len(m.Platforms) > 0 {
		platformSupported := false
		for _, supportedPlatform := range m.Platforms {
			if supportedPlatform == platform || supportedPlatform == "all" {
				platformSupported = true
				break
			}
		}
		if !platformSupported {
			return false
		}
	}
	
	// TODO: Add version comparison logic for CLI version compatibility
	// For now, assume compatible if no version constraints are specified
	if m.CLIVersionMin == "" && m.CLIVersionMax == "" {
		return true
	}
	
	return true // Simplified for initial implementation
}

// GetTierLevel returns the numeric tier level for comparison
func (m *KmpkgMetadata) GetTierLevel() int {
	tierLevels := map[string]int{
		"Free":       0,
		"Pro":        1,
		"Enterprise": 2,
	}
	
	if level, exists := tierLevels[m.RequiredTier]; exists {
		return level
	}
	
	return 0 // Default to Free tier
}