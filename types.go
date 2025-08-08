package kmsdk

import "time"

type Direction string

const (
	DirectionInbound  Direction = "request"
	DirectionOutbound Direction = "response"
)

type SubscriptionTier string

const (
	TierFree       SubscriptionTier = "Free"
	TierPro        SubscriptionTier = "Pro"
	TierEnterprise SubscriptionTier = "Enterprise"
)

type PluginInfo struct {
	Name         string
	Version      string
	Description  string
	RequiredTier SubscriptionTier
}

type Config struct {
	Debug    bool
	Features []string
	// Optional fields that some plugins expect in legacy form
	ApiEndpoint string
}

type StreamEvent struct {
	Type      string
	Timestamp time.Time
	Message   string
}

type Event struct {
	ID        string
	Timestamp string
	Type      string
	Data      map[string]any
}
