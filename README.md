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

## Local development with Go workspaces

By default, consumers fetch a tagged SDK version from the remote. During local development across multiple repos, use a Go workspace to prefer your checked-out sources.

Assuming all repos live under a common parent directory (recommended):

```bash
cd /path/to/parent
go work init ./kilometers-plugins-sdk ./kilometers-cli ./kilometers-cli-plugins
go work use  ./kilometers-plugins-sdk ./kilometers-cli ./kilometers-cli-plugins

# Verify
go env GOWORK
cat go.work
go work sync
```

Notes:
- Keep `go.work` uncommitted (it is a per-developer override).
- VSCode Go extension automatically detects a parent go.work file.
- If your repos are not co-located, you can create a workspace anywhere and `go work use` absolute paths; optionally set `GOWORK=/path/to/go.work`.
