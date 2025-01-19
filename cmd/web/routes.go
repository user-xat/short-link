package main

import "net/http"

// Prescribes the endpoints of the web server
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", app.homeHandler)
	mux.HandleFunc("POST /{$}", app.createShortLinkHandler)
	mux.HandleFunc("GET /{shortlink...}", app.shortLinkHandler)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir(*staticDir)})
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	handler := app.Logging(mux)

	return handler
}
