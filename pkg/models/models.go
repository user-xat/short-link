package models

import "errors"

var ErrNotRecord = errors.New("record has not found")

type LinkData struct {
	Source string `json:"source"`
	Short  string `json:"key"`
}
