package mcp
 import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
 )

 // AddParams defines parameters for addition 
 type AddParams struct {
	A float32 `json:"a" mcp:"first number"`
	B float32 `json:"b" mcp:"second number"`
 }

 //SubtractParams defines parameters for subtraction 
 type SubtractParams struct {
	A float64 `json:"a" mcp:"first number"`
	B float64 `json:"b" mcp:"second number"`
 }

 // MultiplyParams defines parameters for multiplication 
 type MultiplyParams struct {
	A float64 `json:"a" mcp:"first number"`
	B float64 `json:"b" mcp:"second number"`
 }

 // DivideParams defines parameters for division
 type DivideParams struct {
	A float64 `json:"a" mcp:"first number"`
	B float64 `json:"b" mcp:"second number"`
 }

 //Add performs addition of two numbers 
 func Add(_ context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[AddParams]) (*mcp.CallToolResult, error) {
	result := params.Arguments.A + params.Arguments.B
	response := fmt.Sprintf("%.2f + %.2f = %.2f", params.Arguments.A, params.Arguments.B, result)

	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: response}},
	}, nil

 }

// Subtract performs multiplication of two numbers 
func Subtract(_ context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[SubtractParams]) (*mcp.CallToolResult, error) {
	result := params.Arguments.A - params.Arguments.B
	response := fmt.Sprintf("%.2f - %.2f = %.2f", params.Arguments.A, params.Arguments.B, result)

	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: response}},
	}, nil
}

//Multiply performs multiplication of two numbers 
func Multiply(_ context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[MultiplyParams]) (*mcp.CallToolResult, error) {
	result := params.Arguments.A * params.Arguments.B
	response := fmt.Sprintf("%.2f * %.2f = %.2f", params.Arguments.A, params.Arguments.B, result)

	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: response}},
	}, nil
}

// Divide performs division of two numbers 
func Divide(_ context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[DivideParams]) (*mcp.CallToolResult, error) {
	result := params.Arguments.A / params.Arguments.B
	response := fmt.Sprintf("%.2f / %.2f = %.2f", params.Arguments.A, params.Arguments.B, result)

	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: response}},
	}, nil
}

// Tools returns all available tools in MCP 
func Tools() []*mcp.ServerTool {
	tools := []*mcp.ServerTool{
		mcp.NewServerTool(
			"add",
			"Add two numbers",
			Add,
		),

		mcp.NewServerTool(
			"subtract",
			"Subtract two numbers",
			Subtract,
		),

		mcp.NewServerTool(
			"multiply",
			"Multiply two numbers",
			Multiply,
		),

		mcp.NewServerTool(
			"divide",
			"Divide two numbers",
			Divide,
		),
	}

	return  tools
}