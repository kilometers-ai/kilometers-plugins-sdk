# Kilometers Plugins SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/kilometers-ai/kilometers-plugins-sdk.svg)](https://pkg.go.dev/github.com/kilometers-ai/kilometers-plugins-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/kilometers-ai/kilometers-plugins-sdk)](https://goreportcard.com/report/github.com/kilometers-ai/kilometers-plugins-sdk)

Official Go SDK for developing Kilometers CLI plugins with JWT-based authentication and universal binary support.

## Features

- 🔐 **JWT Authentication** - Secure runtime authentication
- 🌍 **Universal Binaries** - Single binary per platform
- 📦 **Simple Interface** - Minimal API surface
- 🔒 **Type Safe** - Full Go type safety
- 📖 **Well Documented** - Complete API documentation

## Installation

```bash
go get github.com/kilometers-ai/kilometers-plugins-sdk
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/kilometers-ai/kilometers-plugins-sdk"
)

type MyPlugin struct {
    authenticated bool
}

// Authenticate validates JWT tokens at runtime
func (p *MyPlugin) Authenticate(ctx context.Context, token string) error {
    // Validate JWT token with your authentication logic
    // Return error if invalid, nil if valid
    p.authenticated = true
    return nil
}

// ProcessMessage handles MCP messages and returns events
func (p *MyPlugin) ProcessMessage(ctx context.Context, message []byte, direction string) ([]kmsdk.Event, error) {
    if !p.authenticated {
        return nil, fmt.Errorf("not authenticated")
    }

    event := kmsdk.Event{
        ID:        fmt.Sprintf("event_%d", time.Now().UnixNano()),
        Timestamp: time.Now().Format(time.RFC3339),
        Type:      "message_processed",
        Data: map[string]interface{}{
            "message_size": len(message),
            "direction":    direction,
        },
    }

    return []kmsdk.Event{event}, nil
}

// GetInfo returns plugin metadata
func (p *MyPlugin) GetInfo() kmsdk.PluginInfo {
    return kmsdk.PluginInfo{
        Name:         "my-plugin",
        Version:      "1.0.0",
        Description:  "My awesome Kilometers plugin",
        RequiredTier: "Free", // Free, Pro, or Enterprise
    }
}
```

## Core Interface

The SDK defines a single interface that all plugins must implement:

```go
type KilometersPlugin interface {
    // Authenticate validates the plugin's JWT token
    Authenticate(ctx context.Context, token string) error
    
    // ProcessMessage handles an MCP message and returns events to log
    ProcessMessage(ctx context.Context, message []byte, direction string) ([]Event, error)
    
    // GetInfo returns plugin metadata
    GetInfo() PluginInfo
}
```

## Types

### Event

Represents a logged event from your plugin:

```go
type Event struct {
    ID        string                 `json:"id"`        // Unique event identifier
    Timestamp string                 `json:"timestamp"` // RFC3339 timestamp
    Type      string                 `json:"type"`      // Event type/category
    Data      map[string]interface{} `json:"data"`      // Event payload
}
```

### PluginInfo

Contains metadata about your plugin:

```go
type PluginInfo struct {
    Name         string `json:"name"`          // Plugin name
    Version      string `json:"version"`       // Plugin version
    Description  string `json:"description"`   // Human-readable description
    RequiredTier string `json:"required_tier"` // Free, Pro, or Enterprise
}
```

## Authentication

Plugins use JWT-based authentication for secure runtime validation:

1. **JWT Validation**: Implement `Authenticate()` to validate JWT tokens
2. **Runtime Checks**: Verify authentication in `ProcessMessage()`
3. **Graceful Failure**: Return errors for invalid authentication

Example JWT validation:

```go
func (p *MyPlugin) Authenticate(ctx context.Context, token string) error {
    // Parse JWT token
    claims, err := jwt.Parse(token, keyFunc)
    if err != nil {
        return fmt.Errorf("invalid JWT: %w", err)
    }
    
    // Validate claims
    if claims.CustomerID == "" {
        return fmt.Errorf("missing customer ID")
    }
    
    // Check tier requirements
    if !p.isValidTier(claims.Tier) {
        return fmt.Errorf("insufficient tier")
    }
    
    p.authenticated = true
    return nil
}
```

## Plugin Distribution

Plugins are distributed as signed `.kmpkg` packages containing:

- Plugin binary
- Manifest with metadata
- Ed25519 signatures for verification

See the [Plugin Distribution Guide](https://docs.kilometers.ai/plugins/distribution) for details.

## Examples

### Console Logger Plugin

```go
func (p *ConsoleLoggerPlugin) ProcessMessage(ctx context.Context, message []byte, direction string) ([]kmsdk.Event, error) {
    // Log to console
    fmt.Printf("[%s] %s: %d bytes\n", time.Now().Format("15:04:05"), direction, len(message))
    
    // Return event for further processing
    return []kmsdk.Event{{
        ID:        generateID(),
        Timestamp: time.Now().Format(time.RFC3339),
        Type:      "console_log",
        Data: map[string]interface{}{
            "message_size": len(message),
            "direction":    direction,
        },
    }}, nil
}
```

### API Logger Plugin

```go
func (p *APILoggerPlugin) ProcessMessage(ctx context.Context, message []byte, direction string) ([]kmsdk.Event, error) {
    event := kmsdk.Event{
        ID:        generateID(),
        Timestamp: time.Now().Format(time.RFC3339),
        Type:      "api_log",
        Data: map[string]interface{}{
            "message":   string(message),
            "direction": direction,
            "customer":  p.customerID,
        },
    }
    
    // Send to API endpoint
    if err := p.sendToAPI(event); err != nil {
        return nil, fmt.Errorf("failed to send to API: %w", err)
    }
    
    return []kmsdk.Event{event}, nil
}
```

## Best Practices

### Security
- Always validate JWT tokens in `Authenticate()`
- Check authentication status in `ProcessMessage()`
- Never log sensitive information (tokens, customer data)
- Fail gracefully with helpful error messages

### Performance
- Keep `ProcessMessage()` fast and non-blocking
- Use buffering for network operations
- Implement proper context cancellation
- Avoid heavy computations in the main path

### Error Handling
- Return descriptive errors from `Authenticate()`
- Handle network failures gracefully
- Use structured logging for debugging
- Implement retry logic for transient failures

### Testing
```go
func TestMyPlugin(t *testing.T) {
    plugin := &MyPlugin{}
    
    // Test authentication
    err := plugin.Authenticate(context.Background(), validToken)
    assert.NoError(t, err)
    
    // Test message processing
    events, err := plugin.ProcessMessage(context.Background(), []byte("test"), "inbound")
    assert.NoError(t, err)
    assert.Len(t, events, 1)
}
```

## Requirements

- Go 1.21 or later
- Valid Kilometers CLI subscription for plugin distribution

## Contributing

This SDK is maintained by the Kilometers team. For questions or issues:

- 📧 Email: support@kilometers.ai
- 📖 Documentation: https://docs.kilometers.ai/plugins/
- 🐛 Issues: Use your private support channel

## License

Proprietary - See LICENSE file for details.

---

Made with ❤️ by the [Kilometers](https://kilometers.ai) team.