package web

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/user-xat/short-link/configs"
	"github.com/user-xat/short-link/pkg/res"
	"github.com/user-xat/short-link/pkg/templates"
)

type WebHandlerDeps struct {
	*WebService
	*configs.WebConfig
	templateCache templates.TemplatesCache
}

type WebHandler struct {
	*WebService
	templateCache templates.TemplatesCache
}

func NewWebHandler(router *http.ServeMux, deps WebHandlerDeps) {
	handler := &WebHandler{
		WebService:    deps.WebService,
		templateCache: deps.templateCache,
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
		url := r.Form.Get("url")
		link, err := h.WebService.CreateLink(url)
		if err != nil {
			res.ServerError(w, h.ErrorLog, err)
			return
		}
		td := &TemplLinkData{
			Url:    url,
			Hashed: fmt.Sprintf("http://%s/%s", r.Host, link),
		}
		err = templates.Render(h.templateCache, w, "home.page.tmpl", td)
		if err != nil {
			res.ServerError(w, h.ErrorLog, err)
		}
	}
}

// func (h *WebHandler) GoTo() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		shortlink := r.PathValue("hash")
// 		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
// 		defer cancel()
// 		if link, err := h.cacheDb.Get(ctx, shortlink); err == nil {
// 			http.Redirect(w, r, link.Source, http.StatusSeeOther)
// 			return
// 		}
// 		link, err := h.serviceClient.Get(ctx, &wrapperspb.StringValue{Value: shortlink})
// 		if err != nil {
// 			if e, ok := status.FromError(err); ok {
// 				switch e.Code() {
// 				case codes.NotFound:
// 					res.ClientError(w, http.StatusNotFound)
// 				default:
// 					res.ServerError(w, h.errorLog, err)
// 				}
// 			} else {
// 				res.ServerError(w, h.errorLog, err)
// 			}
// 			return
// 		}
// 		_, err = h.cacheDb.Set(context.Background(), &models.LinkData{
// 			Short:  link.Short,
// 			Source: link.Source,
// 		})
// 		if err != nil {
// 			h.errorLog.Printf("failed save value to cache: %v", err)
// 		}
// 		http.Redirect(w, r, link.Source, http.StatusSeeOther)
// 	}
// }

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
