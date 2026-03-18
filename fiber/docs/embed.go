package docs

import (
	"embed"
	"io/fs"

	"github.com/betterde/template/fiber/internal/journal"
)

//go:embed api/*
var docs embed.FS

func Serve() fs.FS {
	dist, err := fs.Sub(docs, "orbit")
	if err != nil {
		journal.Logger.Panicw("Error mounting front-end static resources!", err)
	}

	return dist
}
