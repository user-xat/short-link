package req

import (
	"io"
	"net/http"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := GetData[T](r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	return body, nil
}

func GetData[T any](body io.ReadCloser) (*T, error) {
	data, err := Decode[T](body)
	if err != nil {
		return nil, err
	}
	err = IsValid(data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
