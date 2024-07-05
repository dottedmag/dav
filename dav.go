package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/emersion/go-webdav"
)

func main() {
	var rootDir, addr string
	flag.StringVar(&rootDir, "root-dir", "/nonexistent", "Root directory for WebDAV server")
	flag.StringVar(&addr, "addr", "127.0.0.1:6000", "Address to listen to")
	flag.Parse()

	http.Handle("/", &webdav.Handler{
		FileSystem: webdav.LocalFileSystem(rootDir),
	})
	log.Printf("Serving WebDAV %s on %s", rootDir, addr)
	http.ListenAndServe(addr, nil)
}
