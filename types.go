package kmsdk

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