package link_test

import (
	"context"
	"testing"

	"github.com/user-xat/short-link/internal/link"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestRWData(t *testing.T) {
	sl := link.NewGRPCHandler(link.GRPCHandlerDeps{
		LinkRepository: NewLinksStoreMap(),
	})

	rawLink := "http://github.com/"
	link, err := sl.Create(context.Background(), &wrapperspb.StringValue{
		Value: rawLink,
	})
	if err != nil {
		t.Fatalf("can't set value %s", rawLink)
	}
	if link.Url != rawLink {
		t.Fatalf("got %s, want %s", link.Url, rawLink)
	}
	if link.Hash == "" {
		t.Fatal("hash is empty")
	}

	gotLD, err := sl.GetByHash(context.Background(), &wrapperspb.StringValue{
		Value: link.Hash,
	})
	if err != nil {
		t.Fatalf("error to get link by short: %v", err)
	}
	if gotLD.Url != rawLink {
		t.Fatalf("got source link %s, want %s", gotLD.Url, rawLink)
	}
	if gotLD.Hash != link.Hash {
		t.Fatalf("got short link %s, want %s", gotLD.Hash, link.Hash)
	}
}
