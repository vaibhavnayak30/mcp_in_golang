package mcp

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// New server creates and configures new server for math operation 
func NewServer() *mcp.Server {
	server := mcp.NewServer(
		"Math MCP Server",
		"1.0.0",
		&mcp.ServerOptions{
			Instructions: "This MCP Server provides basic mathematical operation: addition, subtraction, multiplication and devision",
		},
	)

	// Add all available tools to serer
	server.AddTools(Tools()...)

	return  server
}