package web

import (
	"log"

	"github.com/user-xat/short-link/configs"
	"github.com/user-xat/short-link/pkg/models/memcached"
)

type WebServiceDeps struct {
	WebConfig *configs.WebConfig
	ErrorLog  *log.Logger
	InfoLog   *log.Logger
	CacheDb   *memcached.CachedLinkModel
}

type WebService struct {
	WebConfig *configs.WebConfig
	ErrorLog  *log.Logger
	InfoLog   *log.Logger
	CacheDb   *memcached.CachedLinkModel
}

func NewWebService(deps WebServiceDeps) *WebService {
	return &WebService{
		ErrorLog:  deps.ErrorLog,
		InfoLog:   deps.InfoLog,
		CacheDb:   deps.CacheDb,
		WebConfig: deps.WebConfig,
	}
}

func (s *WebService) CreateLink(url string) (*Link, error) {
	// ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	// defer cancel()
	// link, err := h.serviceClient.Add(ctx, &wrapperspb.StringValue{Value: sourceLink})
	// if err != nil {
	// 	res.ServerError(w, h.errorLog, fmt.Errorf("remote service: %v", err))
	// 	return
	// }
	return nil, nil
}
