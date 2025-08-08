package kmsdk

import "context"

// LegacyPlugin provides backward compatibility with the old plugin interface
// This allows existing plugins to work with minimal changes during migration
type LegacyPlugin interface {
	// Metadata
	Name() string
	RequiredFeature() string
	RequiredTier() SubscriptionTier

	// Lifecycle - similar to SDK but with different dependencies structure
	Initialize(ctx context.Context, deps LegacyPluginDependencies) error
	Shutdown(ctx context.Context) error

	// Message handling
	HandleMessage(ctx context.Context, data []byte, direction Direction) error
	HandleError(ctx context.Context, err error)
	HandleStreamEvent(ctx context.Context, event StreamEvent)
}

// LegacyPluginDependencies represents the dependencies passed to legacy plugins
type LegacyPluginDependencies struct {
	Config      Config
	AuthManager LegacyAuthManager
}

// LegacyAuthManager provides authentication and feature checking
type LegacyAuthManager interface {
	IsFeatureEnabled(feature string) bool
	GetSubscriptionTier() SubscriptionTier
}

// LegacyToSDKAdapter adapts a LegacyPlugin to implement the SDK Plugin interface
type LegacyToSDKAdapter struct {
	legacy LegacyPlugin
	deps   LegacyPluginDependencies
}

// NewLegacyAdapter creates an adapter to wrap legacy plugins
func NewLegacyAdapter(legacy LegacyPlugin) *LegacyToSDKAdapter {
	return &LegacyToSDKAdapter{legacy: legacy}
}

// Authenticate implements SDK Plugin interface
func (a *LegacyToSDKAdapter) Authenticate(ctx context.Context, token string) error {
	// Legacy plugins don't have separate authentication step
	// Authentication is handled during Initialize
	return nil
}

// Initialize implements SDK Plugin interface
func (a *LegacyToSDKAdapter) Initialize(ctx context.Context, cfg Config) error {
	// Convert SDK config to legacy dependencies
	a.deps = LegacyPluginDependencies{
		Config: cfg,
		AuthManager: &simpleLegacyAuthManager{
			features: cfg.Features,
		},
	}
	return a.legacy.Initialize(ctx, a.deps)
}

// Shutdown implements SDK Plugin interface
func (a *LegacyToSDKAdapter) Shutdown(ctx context.Context) error {
	return a.legacy.Shutdown(ctx)
}

// HandleMessage implements SDK Plugin interface
func (a *LegacyToSDKAdapter) HandleMessage(ctx context.Context, data []byte, direction Direction, correlationID string) error {
	return a.legacy.HandleMessage(ctx, data, direction)
}

// HandleError implements SDK Plugin interface
func (a *LegacyToSDKAdapter) HandleError(ctx context.Context, err error) {
	a.legacy.HandleError(ctx, err)
}

// HandleStreamEvent implements SDK Plugin interface
func (a *LegacyToSDKAdapter) HandleStreamEvent(ctx context.Context, event StreamEvent) {
	a.legacy.HandleStreamEvent(ctx, event)
}

// GetInfo implements SDK Plugin interface
func (a *LegacyToSDKAdapter) GetInfo() PluginInfo {
	return PluginInfo{
		Name:         a.legacy.Name(),
		Version:      "1.0.0", // Default version for legacy plugins
		Description:  "Legacy plugin: " + a.legacy.Name(),
		RequiredTier: a.legacy.RequiredTier(),
	}
}

// simpleLegacyAuthManager provides a simple implementation for legacy compatibility
type simpleLegacyAuthManager struct {
	features []string
}

func (m *simpleLegacyAuthManager) IsFeatureEnabled(feature string) bool {
	for _, f := range m.features {
		if f == feature {
			return true
		}
	}
	return false
}

func (m *simpleLegacyAuthManager) GetSubscriptionTier() SubscriptionTier {
	// This would need to be determined from the features list or config
	// For now, assume Pro if API logging is available, otherwise Free
	if m.IsFeatureEnabled(FeatureAPILogging) {
		return TierPro
	}
	return TierFree
}