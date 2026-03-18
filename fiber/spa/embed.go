package spa

import (
	"embed"
	"io/fs"

	"github.com/betterde/template/fiber/internal/journal"
)

//go:embed dist/*
var spa embed.FS

func Serve() fs.FS {
	dist, err := fs.Sub(spa, "dist")
	if err != nil {
		journal.Logger.Panicw("Error mounting front-end static resources!", err)
	}

	return dist
}
