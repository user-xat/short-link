package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", app.homeHandler)
	mux.HandleFunc("POST /create", app.createShortLinkHandler)
	mux.HandleFunc("GET /{shortlink}", app.shortLinkHandler)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir(*staticDir)})
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
