package memcached

import (
	"context"
	"testing"
	"time"

	"github.com/user-xat/short-link-server/pkg/models"
)

func TestRWData(t *testing.T) {
	m := NewCachedLinkModel("localhost:11211")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data := &models.LinkData{
		Source: "http://github.com/",
		Short:  "code1",
	}

	// insert data
	key, err := m.Set(ctx, data)
	if err != nil {
		t.Error(err)
	}
	if key != data.Short {
		t.Errorf("got %s, want %s", key, data.Short)
	}

	// get data
	link, err := m.Get(ctx, key)
	if err != nil {
		t.Error(err)
	}
	if link.Short != data.Short || link.Source != data.Source {
		t.Errorf("got link.Source=%s and link.Short=%s, want link.Source=%s and link.Short=%s",
			link.Source, link.Short, data.Source, data.Short)
	}
}
