package kmsdk

import (
	"context"
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// HandshakeConfig is the shared handshake config for Kilometers plugins
var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "KILOMETERS_PLUGIN",
	MagicCookieValue: "kilometers_monitoring_plugin",
}

// KilometersPluginRPC implements plugin.Plugin for net/rpc
type KilometersPluginRPC struct {
	// Impl is the concrete implementation, only used on the plugin side
	Impl KilometersPlugin
}

// Server returns an RPC server for this plugin
func (p *KilometersPluginRPC) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RPCServer{Impl: p.Impl}, nil
}

// Client returns an RPC client for this plugin
func (p *KilometersPluginRPC) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RPCClient{client: c}, nil
}

// KilometersPluginGRPC implements plugin.Plugin for gRPC (for future use if needed)
type KilometersPluginGRPC struct {
	plugin.NetRPCUnsupportedPlugin
	// Impl is the concrete implementation, only used on the plugin side
	Impl KilometersPlugin
}

// GRPCServer is currently not implemented - we use net/rpc
func (p *KilometersPluginGRPC) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	// Not implemented - using net/rpc instead
	return nil
}

// GRPCClient is currently not implemented - we use net/rpc
func (p *KilometersPluginGRPC) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	// Not implemented - using net/rpc instead
	return nil, nil
}

// PluginMap returns the default plugin map for Kilometers plugins
// This should be used by both the host and plugin sides
func PluginMap(impl KilometersPlugin) map[string]plugin.Plugin {
	return map[string]plugin.Plugin{
		"kilometers": &KilometersPluginRPC{Impl: impl},
	}
}

// ServePlugin is a helper function for plugins to serve themselves
func ServePlugin(impl KilometersPlugin) {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: HandshakeConfig,
		Plugins:         PluginMap(impl),
	})
}