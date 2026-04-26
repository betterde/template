package example

import "github.com/modelcontextprotocol/go-sdk/mcp"

func Register(server *mcp.Server) {
	mcp.AddTool(server, echoTool, echoHandler)
}
