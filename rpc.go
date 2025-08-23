package kmsdk

import (
	"context"
	"net/rpc"
)

// RPCClient is the RPC client implementation of KilometersPlugin
type RPCClient struct {
	client *rpc.Client
}

// Authenticate performs runtime JWT validation
func (c *RPCClient) Authenticate(ctx context.Context, token string) error {
	var resp error
	err := c.client.Call("Plugin.Authenticate", token, &resp)
	if err != nil {
		return err
	}
	return resp
}

// ProcessMessage handles an MCP message and returns events to log
func (c *RPCClient) ProcessMessage(ctx context.Context, message []byte, direction string) ([]Event, error) {
	args := ProcessMessageArgs{
		Message:   message,
		Direction: direction,
	}
	var resp ProcessMessageResponse
	err := c.client.Call("Plugin.ProcessMessage", args, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Error != "" {
		return nil, NewPluginError(resp.Error)
	}
	return resp.Events, nil
}

// GetInfo returns plugin metadata
func (c *RPCClient) GetInfo() PluginInfo {
	var resp PluginInfo
	err := c.client.Call("Plugin.GetInfo", new(interface{}), &resp)
	if err != nil {
		// Return default info on error
		return PluginInfo{
			Name:    "unknown",
			Version: "unknown",
		}
	}
	return resp
}

// RPCServer is the RPC server wrapper for a KilometersPlugin implementation
type RPCServer struct {
	Impl KilometersPlugin
}

// Authenticate is the RPC server method
func (s *RPCServer) Authenticate(token string, resp *error) error {
	*resp = s.Impl.Authenticate(context.Background(), token)
	return nil
}

// ProcessMessage is the RPC server method
func (s *RPCServer) ProcessMessage(args ProcessMessageArgs, resp *ProcessMessageResponse) error {
	events, err := s.Impl.ProcessMessage(context.Background(), args.Message, args.Direction)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}
	resp.Events = events
	return nil
}

// GetInfo is the RPC server method
func (s *RPCServer) GetInfo(args interface{}, resp *PluginInfo) error {
	*resp = s.Impl.GetInfo()
	return nil
}

// ProcessMessageArgs holds the arguments for ProcessMessage RPC
type ProcessMessageArgs struct {
	Message   []byte
	Direction string
}

// ProcessMessageResponse holds the response for ProcessMessage RPC
type ProcessMessageResponse struct {
	Events []Event
	Error  string
}

