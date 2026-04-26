package tools

import (
	"github.com/betterde/template/mcp/tools/example"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Register(server *mcp.Server) {
	example.Register(server)
}
