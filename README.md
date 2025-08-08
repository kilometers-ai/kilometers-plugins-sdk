# Kilometers Plugin SDK

Minimal, stable contract for Kilometers CLI plugins.

## Import

- `go get github.com/kilometers-ai/kilometers-plugins-sdk`
- `import "github.com/kilometers-ai/kilometers-plugins-sdk"`

## Versioning

- Start with v0.x during stabilization.
- Move to v1.0.0 once contracts are stable; follow SemVer for breaking changes.

## Package

- Root package `kmsdk` exposes:
  - `Plugin` interface
  - `Config`, `PluginInfo`, `StreamEvent`, `Event`
  - `Direction`, `SubscriptionTier` enums
