package web

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type LinkCreateResponse struct {
	Id   int    `json:"ID"`
	Url  string `json:"url"`
	Hash string `json:"hash"`
}

type TemplLinkData struct {
	Url    string
	Hashed string
}
