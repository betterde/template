package docs

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/betterde/template/fiber/internal/journal"
)

//go:embed api/*
var FS embed.FS

func Serve() http.FileSystem {
	dist, err := fs.Sub(FS, "orbit")
	if err != nil {
		journal.Logger.Panicw("Error mounting front-end static resources!", err)
	}

	return http.FS(dist)
}
