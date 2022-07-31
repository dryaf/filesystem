package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed files
var embededFiles embed.FS

func getFileSystem(webFiles string, dontUseEmbededFiles bool) http.FileSystem {
	if dontUseEmbededFiles {
		log.Print("using local file mode")
		return http.FS(os.DirFS(webFiles))
	}

	log.Print("using embed mode")
	fsys, err := fs.Sub(embededFiles, webFiles)
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
