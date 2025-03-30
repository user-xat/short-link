package link

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/user-xat/short-link/configs"
	"github.com/user-xat/short-link/internal/models"
	"github.com/user-xat/short-link/pkg/event"
	"github.com/user-xat/short-link/pkg/middleware"
	"github.com/user-xat/short-link/pkg/req"
	"github.com/user-xat/short-link/pkg/res"
	pb "github.com/user-xat/short-link/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gorm.io/gorm"
)

type ICache interface {
	Get(context.Context, string) (*models.Link, error)
	Set(context.Context, *models.Link) error
}

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
	Config         *configs.ApiConfig
	EventBus       *event.EventBus
	Service        pb.ShortLinkClient
	ErrorLog       *log.Logger
	InfoLog        *log.Logger
	Cache          ICache
}

type LinkHandler struct {
	LinkRepository *LinkRepository
	EventBus       *event.EventBus
	Service        pb.ShortLinkClient
	errorLog       *log.Logger
	infoLog        *log.Logger
	cache          ICache
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
		EventBus:       deps.EventBus,
		Service:        deps.Service,
		infoLog:        deps.InfoLog,
		errorLog:       deps.ErrorLog,
		cache:          deps.Cache,
	}

	router.HandleFunc("GET /{hash}", handler.GoTo())
	router.HandleFunc("POST /link", handler.Create())
	router.Handle("GET /link", middleware.IsAuthed(handler.GetAll(), deps.Config))
	router.Handle("PATCH /link/{id}", middleware.IsAuthed(handler.Update(), deps.Config))
	router.Handle("DELETE /link/{id}", middleware.IsAuthed(handler.Delete(), deps.Config))
}

func (h *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkCreateRequest](w, r)
		if err != nil {
			return
		}
		createdLink, err := h.Service.Create(r.Context(), wrapperspb.String(body.Url))
		if err != nil {
			res.ServerError(w, h.errorLog, err)
			return
		}
		res.Json(w, models.Link{
			Model: gorm.Model{ID: uint(createdLink.Id)},
			Url:   createdLink.Url,
			Hash:  createdLink.Hash,
		}, http.StatusCreated)
	}
}

func (h *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		if link, err := h.cache.Get(r.Context(), hash); err == nil {
			http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
			return
		}
		link, err := h.Service.GetByHash(r.Context(), wrapperspb.String(hash))
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
		go h.EventBus.Publish(event.Event{
			Type: event.EventLinkVisited,
			Data: &models.Link{
				Model: gorm.Model{ID: uint(link.GetId())},
				Url:   link.GetUrl(),
				Hash:  link.GetHash(),
			},
		})
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}

func (h *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email, ok := r.Context().Value(middleware.ContextEmailKey).(string)
		if ok {
			fmt.Println(email)
		}
		body, err := req.HandleBody[LinkUpdateRequest](w, r)
		if err != nil {
			return
		}

		sId := r.PathValue("id")
		id, err := strconv.ParseUint(sId, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		link, err := h.Service.Update(r.Context(), &pb.Link{
			Id:   id,
			Url:  body.Url,
			Hash: body.Hash,
		})
		if err != nil {
			res.ServerError(w, h.errorLog, err)
			return
		}
		updatedLink := models.Link{
			Model: gorm.Model{ID: uint(link.Id)},
			Url:   link.Url,
			Hash:  link.Hash,
		}
		go h.EventBus.Publish(event.Event{
			Type: event.EventLinkUpdated,
			Data: &updatedLink,
		})
		res.Json(w, updatedLink, http.StatusOK)
	}
}

func (h *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sId := r.PathValue("id")
		id, err := strconv.ParseUint(sId, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = h.Service.Delete(r.Context(), wrapperspb.UInt64(id))
		if err != nil {
			res.ServerError(w, h.errorLog, err)
			return
		}
		res.Json(w, nil, http.StatusOK)
	}
}

func (h *LinkHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			http.Error(w, "invalid limit", http.StatusBadRequest)
			return
		}
		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			http.Error(w, "invalid offset", http.StatusBadRequest)
			return
		}
		recLinks, err := h.Service.GetAll(r.Context(), &pb.LimitOffset{
			Limit:  uint64(limit),
			Offset: uint64(offset),
		})
		if err != nil {
			res.ServerError(w, h.errorLog, err)
			return
		}
		recCount, err := h.Service.Count(r.Context(), &pb.Void{})
		if err != nil {
			res.ServerError(w, h.errorLog, err)
			return
		}
		pbLinks := recLinks.GetLink()
		links := make([]models.Link, len(pbLinks))
		for i := range links {
			links[i] = models.Link{
				Model: gorm.Model{ID: uint(pbLinks[i].Id)},
				Url:   pbLinks[i].Url,
				Hash:  pbLinks[i].Hash,
			}
		}
		res.Json(w, GetAllLinksResponse{
			Links: links,
			Count: int64(recCount.GetValue()),
		}, http.StatusOK)
	}
}
