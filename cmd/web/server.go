package main

import (
	"html/template"
	"log"
	"net/http"
)

type Server struct {
	shortLink *ShortLink
}

func NewServer(store LinksStore) *Server {
	return &Server{
		shortLink: NewShortLink(store),
	}
}

func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
