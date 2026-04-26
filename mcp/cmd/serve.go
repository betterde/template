/*
Copyright © 2026 George <george@betterde.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/betterde/template/mcp/config"
	"github.com/betterde/template/mcp/global"
	"github.com/betterde/template/mcp/internal/journal"
	"github.com/betterde/template/mcp/internal/middleware"
	"github.com/betterde/template/mcp/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the MCP server",
	Run: func(cmd *cobra.Command, args []string) {
		start(global.Ctx)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func start(ctx context.Context) {
	server := mcp.NewServer(&mcp.Implementation{
		Name:       "mysql-mcp-server",
		Version:    "v1.0.0",
		WebsiteURL: "https://github.com/betterde/mysql-mcp-server",
	}, nil)

	server.AddReceivingMiddleware(middleware.Logging)

	tools.Register(server)

	loggerHandler := zapslog.NewHandler(journal.Logger.Core(), zapslog.WithName(journal.Logger.Name()))

	handler := mcp.NewStreamableHTTPHandler(func(req *http.Request) *mcp.Server {
		return server
	}, &mcp.StreamableHTTPOptions{
		Logger:                     slog.New(loggerHandler),
		Stateless:                  true,
		DisableLocalhostProtection: true,
	})

	mux := http.NewServeMux()
	mux.Handle("/", handler)

	journal.Logger.Info("Starting MCP streamable HTTP server",
		zap.String("addr", config.Conf.HTTP.Listen),
		zap.String("endpoint", "/"),
	)

	go func() {
		if err := http.ListenAndServe(config.Conf.HTTP.Listen, mux); err != nil {
			journal.Logger.Panic(err.Error())
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	<-signalChan
	fmt.Print("\r\033[K")
	journal.Logger.Sugar().Info("Shutdown signal received, exiting...")

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
}
