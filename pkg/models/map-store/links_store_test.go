package mapstore

import (
	"context"
	"testing"

	"github.com/user-xat/short-link-server/pkg/models"
)

func TestRWData(t *testing.T) {
	store := NewLinksStoreMap()

	link := &models.LinkData{
		Source: "http://vk.com/",
		Short:  "AbCd18",
	}
	// insert data
	key, err := store.Set(context.Background(), link)
	if err != nil {
		t.Errorf("failed to set data, error: %s\n", err.Error())
	}

	// read data
	res, err := store.Get(context.Background(), key)
	if err != nil {
		t.Errorf("failed to get value, error: %v\n", err.Error())
	}
	if res.Source != link.Source {
		t.Errorf("got value=%s, want %s", res.Source, link.Source)
	}
}
