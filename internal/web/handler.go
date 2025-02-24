package web

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/user-xat/short-link/configs"
	"github.com/user-xat/short-link/pkg/models"
	"github.com/user-xat/short-link/pkg/res"
	"github.com/user-xat/short-link/pkg/templates"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type WebHandlerDeps struct {
	*WebService
	*configs.ApiConfig
}

type WebHandler struct {
	*WebService
}

func NewWebHandler(router *http.ServeMux, deps WebHandlerDeps) {
	handler := &WebHandler{
		WebService: deps.WebService,
	}

	router.HandleFunc("GET /{$}", handler.Home())
	router.HandleFunc("POST /{$}", handler.CreateShortLink())
	router.HandleFunc("GET /{hash}", handler.GoTo())

	fileServer := http.FileServer(neuteredFileSystem{http.Dir(*staticDir)})
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

		sourceLink := r.Form.Get("url")

		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()
		link, err := h.serviceClient.Add(ctx, &wrapperspb.StringValue{Value: sourceLink})
		if err != nil {
			res.ServerError(w, h.errorLog, fmt.Errorf("remote service: %v", err))
			return
		}

		td := &templates.TemplateData{Link: &models.LinkData{
			Source: link.Source,
			Short:  fmt.Sprintf("http://%s/%s", r.Host, link.Short),
		}}

		err = templates.Render(h.templateCache, w, "home.page.tmpl", td)
		if err != nil {
			res.ServerError(w, h.errorLog, err)
		}
	}
}

func (h *WebHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortlink := r.PathValue("shortlink")
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		if link, err := h.cacheDb.Get(ctx, shortlink); err == nil {
			http.Redirect(w, r, link.Source, http.StatusSeeOther)
			return
		}

		link, err := h.serviceClient.Get(ctx, &wrapperspb.StringValue{Value: shortlink})
		if err != nil {
			if e, ok := status.FromError(err); ok {
				switch e.Code() {
				case codes.NotFound:
					res.ClientError(w, http.StatusNotFound)
				default:
					res.ServerError(w, h.errorLog, err)
				}
			} else {
				res.ServerError(w, h.errorLog, err)
			}
			return
		}

		_, err = h.cacheDb.Set(context.Background(), &models.LinkData{
			Short:  link.Short,
			Source: link.Source,
		})
		if err != nil {
			h.errorLog.Printf("failed save value to cache: %v", err)
		}

		http.Redirect(w, r, link.Source, http.StatusSeeOther)
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
