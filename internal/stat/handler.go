package stat

import (
	"net/http"
	"time"

	"github.com/user-xat/short-link/configs"
	"github.com/user-xat/short-link/pkg/middleware"
	"github.com/user-xat/short-link/pkg/res"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.ApiConfig
}

type StatHandler struct {
	StatRepository *StatRepository
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
	}
	router.Handle("GET /stat", middleware.IsAuthed(handler.Get(), deps.Config))
}

func (h *StatHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, err := time.Parse(time.DateOnly, r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "invalid from param", http.StatusBadRequest)
			return
		}
		to, err := time.Parse(time.DateOnly, r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "invalid to param", http.StatusBadRequest)
			return
		}
		by := r.URL.Query().Get("by")
		if by != GroupByDay && by != GroupByMonth {
			http.Error(w, "invalid by param", http.StatusBadRequest)
			return
		}
		stats := h.StatRepository.Get(by, from, to)
		res.Json(w, GetStatsResponse{
			Stats: stats,
		}, http.StatusOK)
	}
}
