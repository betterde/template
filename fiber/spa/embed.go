package spa

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/betterde/template/fiber/internal/journal"
)

//go:embed dist/*
var FS embed.FS

func Serve() http.FileSystem {
	dist, err := fs.Sub(FS, "dist")
	if err != nil {
		journal.Logger.Panicw("Error mounting front-end static resources!", err)
	}

	return http.FS(dist)
}
