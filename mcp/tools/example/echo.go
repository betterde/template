package example

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var echoTool = &mcp.Tool{
	Name:        "echo",
	Title:       "Echo",
	Description: "This is an example of echo",
}

type EchoInput struct {
	Content string `json:"content"`
}

type EchoOutput struct {
	Content string `json:"content"`
}

func echoHandler(ctx context.Context, req *mcp.CallToolRequest, input EchoInput) (*mcp.CallToolResult, EchoOutput, error) {
	output := EchoOutput{
		Content: input.Content,
	}
	return &mcp.CallToolResult{}, output, nil
}
