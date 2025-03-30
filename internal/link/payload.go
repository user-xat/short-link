package link

import "github.com/user-xat/short-link/internal/models"

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type LinkUpdateRequest struct {
	Url  string `json:"url" validate:"required,url"`
	Hash string `json:"hash"`
}

type GetAllLinksResponse struct {
	Links []models.Link `json:"links"`
	Count int64         `json:"count"`
}
