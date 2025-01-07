package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/user-xat/short-link-server/pkg/models/memcached"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	cache    *memcached.CachedLinkModel
	// shortLink *ShortLink
}

func NewServer(errorLog, infoLog *log.Logger, cacheServers []string) *application {
	return &application{
		// shortLink: NewShortLink(store),
		errorLog: errorLog,
		infoLog:  infoLog,
		cache:    memcached.NewCachedLinkModel(cacheServers...),
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

func (app *application) createShortLinkHandler(w http.ResponseWriter, r *http.Request) {
	// TODO:
}

func (app *application) shortLinkHandler(w http.ResponseWriter, r *http.Request) {
	shortlink := r.PathValue("shortlink")

	if link, err := app.cache.Get(shortlink); err == nil {
		http.Redirect(w, r, link.Source, http.StatusSeeOther)
		return
	}
}
