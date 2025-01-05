package main

import (
	"html/template"
	"log"
	"net/http"
)

type application struct {
	errorLog  *log.Logger
	infoLog   *log.Logger
	shortLink *ShortLink
}

func NewServer(store LinksStore, errorLog, infoLog *log.Logger) *application {
	return &application{
		shortLink: NewShortLink(store),
		errorLog:  errorLog,
		infoLog:   infoLog,
	}
}

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}
