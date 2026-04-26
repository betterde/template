package middleware

import (
	"context"
	"time"

	"github.com/betterde/template/mcp/internal/journal"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"go.uber.org/zap"
)

func Logging(next mcp.MethodHandler) mcp.MethodHandler {
	return func(ctx context.Context, method string, req mcp.Request) (result mcp.Result, err error) {
		journal.Logger.Debug(
			"MCP method started",
			zap.String("method", method),
			zap.String("session_id", req.GetSession().ID()),
		)

		// Log more for tool calls.
		if ctr, ok := req.(*mcp.CallToolRequest); ok {
			journal.Logger.Debug("Calling tool",
				zap.String("method", method),
				zap.String("name", ctr.Params.Name),
				zap.Any("args", ctr.Params.Arguments),
			)
		}

		start := time.Now()
		result, err = next(ctx, method, req)
		duration := time.Since(start)
		if err != nil {
			journal.Logger.Error("MCP method failed",
				zap.String("method", method),
				zap.String("session_id", req.GetSession().ID()),
				zap.Int64("duration_ms", duration.Milliseconds()),
				zap.NamedError("err", err),
			)
		} else {
			if ctr, ok := result.(*mcp.CallToolResult); ok {
				journal.Logger.Debug("Result",
					zap.Any("structuredContent", ctr.StructuredContent),
				)
			}
		}
		return result, err
	}
}
