package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	port              = flag.String("addr", "8110", "Launch port")
	memCacheAddr      = flag.String("memcache", "cache:11211", "Network addres for Memcached")
	remoteServiceAddr = flag.String("remote-service", "service:54321", "The addres remote service")
	staticDir         = flag.String("static-dir", "./ui/static", "Path to static assets")
	htmlTemplates     = flag.String("html-templ-dir", "./ui/html", "path to html templates dir")
)

func main() {
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app, err := NewApplication(errorLog, infoLog, *htmlTemplates, *remoteServiceAddr, []string{*memCacheAddr})
	if err != nil {
		errorLog.Fatalf("failed create application: %v", err)
	}
	defer app.Close()

	srv := &http.Server{
		Addr:     ":" + *port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Server launch on http://localhost:%v", *port)
	errorLog.Fatal(srv.ListenAndServe())
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
