package web

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/user-xat/short-link/configs"
	"github.com/user-xat/short-link/pkg/req"
	"github.com/user-xat/short-link/pkg/res"
	"github.com/user-xat/short-link/pkg/templates"
)

type WebHandlerDeps struct {
	*WebService
	*configs.WebConfig
	TemplateCache templates.TemplatesCache
}

type WebHandler struct {
	*WebService
	templateCache templates.TemplatesCache
}

func NewWebHandler(router *http.ServeMux, deps WebHandlerDeps) {
	handler := &WebHandler{
		WebService:    deps.WebService,
		templateCache: deps.TemplateCache,
	}

	router.HandleFunc("GET /{$}", handler.Home())
	router.HandleFunc("POST /{$}", handler.CreateShortLink())

	fileServer := http.FileServer(neuteredFileSystem{http.Dir(deps.StaticDir)})
	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))
}

func (h *WebHandler) Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := templates.Render(h.templateCache, w, "home.page.tmpl", nil)
		if err != nil {
			res.ServerError(w, nil, err)
			return
		}
	}
}

func (h *WebHandler) CreateShortLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if !r.Form.Has("url") || r.Form.Get("url") == "" {
			res.ClientError(w, http.StatusBadRequest)
			return
		}
		reqLink := LinkCreateRequest{Url: r.FormValue("url")}
		err := req.IsValid(reqLink)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		link, err := h.WebService.CreateLink(&reqLink)
		if err != nil {
			res.ServerError(w, h.ErrorLog, err)
			return
		}
		td := &TemplLinkData{
			Url:    link.Url,
			Hashed: fmt.Sprintf("%s/%s", h.WebConfig.ApiAddr, link.Hash),
		}
		err = templates.Render(h.templateCache, w, "home.page.tmpl", td)
		if err != nil {
			res.ServerError(w, h.ErrorLog, err)
		}
	}
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
