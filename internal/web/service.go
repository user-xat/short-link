package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/user-xat/short-link/configs"
	"github.com/user-xat/short-link/pkg/req"
)

const (
	ReqCreateLink = "CreateLink"
)

type WebServiceDeps struct {
	WebConfig *configs.WebConfig
	ErrorLog  *log.Logger
	InfoLog   *log.Logger
}

type WebService struct {
	WebConfig *configs.WebConfig
	ErrorLog  *log.Logger
	InfoLog   *log.Logger
	client    *http.Client
}

func NewWebService(deps WebServiceDeps) *WebService {
	return &WebService{
		ErrorLog:  deps.ErrorLog,
		InfoLog:   deps.InfoLog,
		WebConfig: deps.WebConfig,
		client:    &http.Client{},
	}
}

func (s *WebService) CreateLink(link *LinkCreateRequest) (*LinkCreateResponse, error) {
	r, err := s.createRequest(ReqCreateLink, link)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	linkRes, err := req.Decode[LinkCreateResponse](resp.Body)
	return &linkRes, err
}

func (s *WebService) createRequest(reqType string, data any) (*http.Request, error) {
	var method string
	var body io.Reader
	u, err := url.Parse(s.WebConfig.ApiAddr)
	if err != nil {
		return nil, err
	}
	switch reqType {
	case ReqCreateLink:
		method = http.MethodPost
		u.Path = "/link"
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(jsonData)
	default:
		return nil, errors.New("this type of request is not supported")
	}
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
