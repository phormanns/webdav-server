package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/webdav"
)

func main() {
	dirFlag := flag.String("dir", "", "Directory to serve from. Default: CWD")
	listenAddr := flag.String("addr", "localhost:8080", "IP adress and port to listen on. Default: localhost:8080")
	prefixUrl := flag.String("prefix", "", "URL to strip from resource paths. None by default")
	rootUrl := flag.String("url", "/", "Root url to handle. Default: /")
	flag.Parse()

	dir := *dirFlag

	srv := &webdav.Handler{
		FileSystem: webdav.Dir(dir),
		LockSystem: webdav.NewMemLS(),
		Prefix:     *prefixUrl,
		Logger: func(r *http.Request, err error) {
			if err != nil {
				log.Printf("WEBDAV [%s]: %s, ERROR: %s\n", r.Method, r.URL, err)
			} else {
				log.Printf("WEBDAV [%s]: %s\n", r.Method, r.URL)
			}
		},
	}
	http.Handle(*rootUrl, srv)

	if err := http.ListenAndServe(fmt.Sprintf("%s", *listenAddr), nil); err != nil {
		log.Fatalf("Error with WebDAV server: %v", err)
	}
}
