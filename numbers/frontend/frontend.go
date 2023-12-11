package frontend

import (
	"embed"
	"github.com/gin-contrib/static"
	"net/http"
	"strings"
)

//go:embed *
var embedFrontend embed.FS

type embedFileSystem struct {
	http.FileSystem
	indexes bool
}

const INDEX = "index.html"

func GetFrontendAssets(indexing bool) static.ServeFileSystem {
	return &embedFileSystem{
		FileSystem: http.FS(embedFrontend),
		indexes:    indexing,
	}
}

func (e *embedFileSystem) Exists(prefix string, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		f, err := e.Open(filepath)
		if err != nil {
			return false
		}
		stats, _ := f.Stat()
		if stats.IsDir() {
			if !e.indexes {
				_, err = e.FileSystem.Open(INDEX)
				if err != nil {
					return false
				}
			}
		}
		return true
	}
	return false
}
