package main

import (
	"context"
	"testing"

	mapstore "github.com/user-xat/short-link-server/pkg/models/map-store"
)

func TestRWData(t *testing.T) {
	store := mapstore.NewLinksStoreMap()
	sl := NewShortLink(store)

	rawLink := "http://github.com/"
	link, err := sl.Set(context.Background(), rawLink)
	if err != nil {
		t.Fatalf("can't set value %s", rawLink)
	}
	if link.Source != rawLink {
		t.Fatalf("got %s, want %s", link.Source, rawLink)
	}
	if link.Short == "" {
		t.Fatal("short link is empty")
	}

	gotLD, err := sl.GetLink(context.Background(), link.Short)
	if err != nil {
		t.Fatalf("error to get link by short: %v", err)
	}
	if gotLD.Source != rawLink {
		t.Fatalf("got source link %s, want %s", gotLD.Source, rawLink)
	}
	if gotLD.Short != link.Short {
		t.Fatalf("got short link %s, want %s", gotLD.Short, link.Short)
	}
}
