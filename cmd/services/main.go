// Package main provides entry point to Math MCP server
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	mcpinternal "github.com/vaibhavnayak30/mcp_in_golang/internal/mcp"
	"github.com/vaibhavnayak30/mcp_in_golang/internal/router"
)

// @title MCP Math Server API
// @version 1.0
// @description Math MCP Server for basic math operations

//BasePath
func main() {
	ctx := context.Background()

	//Create MCP server instance 
	server := mcpinternal.NewServer()

	// Create Streamable HTTP handler for JSON-RPC over HTTP 
	streamableHTTPHandler := mcp.NewStreamableHTTPHandler(
		func(_ *http.Request) *mcp.Server {
			return server
		}, nil)
	
	// Creste SSE handler for Server-Sent Events
	sseHandler := mcp.NewSSEHandler(
		func(_ *http.Request) *mcp.Server {
			return server
		}
	)

	// Setup HTTP router
	r := router.SetupRouter()

	// Register MCP endpoints 
	r.HandleFunc("/mcp", func(w http.ResponseWriter, r *http.Request) {
		streamableHTTPHandler.ServeHTTP(w,r)
	})

	r.HandleFunc("/sse", func(w http.ResponseWriter, r *http.Request) {
		sseHandler.ServeHTTP(w,r)
	})

	// Launch the server 
	launch(ctx, r, 8994)
}
