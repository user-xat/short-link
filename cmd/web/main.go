package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	addr              = flag.String("addr", "localhost:8080", "Network address for HTTP")
	staticDir         = flag.String("static-dir", "./ui/static", "Path to static assets")
	memCacheAddr      = flag.String("memcache", "localhost:11211", "Network addres for Memcached")
	remoteServiceAddr = flag.String("remote-service", "localhost:50051", "The addres remote service")
)

func main() {
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app, err := NewApplication(errorLog, infoLog, *remoteServiceAddr, []string{*memCacheAddr})
	if err != nil {
		errorLog.Fatalf("failed create application: %v", err)
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Server launch on http://%v", *addr)
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
