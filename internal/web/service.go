package web

import (
	"log"

	"github.com/user-xat/short-link/pkg/models/memcached"
	"github.com/user-xat/short-link/pkg/templates"
)

type WebService struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	cacheDb       *memcached.CachedLinkModel
	templateCache templates.TemplatesCache
}

func NewWebService() *WebService {
	return &WebService{}
}
