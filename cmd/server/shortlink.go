package main

import (
	"context"
	"crypto/sha1"
	"encoding/base64"

	"github.com/user-xat/short-link-server/pkg/models"
)

type LinksStore interface {
	Get(context.Context, string) (*models.LinkData, error)
	Set(context.Context, *models.LinkData) (string, error)
}

type ShortLink struct {
	store LinksStore
}

func NewShortLink(store LinksStore) *ShortLink {
	return &ShortLink{store: store}
}

func (sl *ShortLink) GetLink(ctx context.Context, suffix string) (*models.LinkData, error) {
	return sl.store.Get(ctx, suffix)
}

func (sl *ShortLink) Set(ctx context.Context, link string) (*models.LinkData, error) {
	suffix := generateSuffix(link)
	_, err := sl.store.Set(ctx, &models.LinkData{
		Short:  suffix,
		Source: link,
	})
	if err != nil {
		return nil, err
	}

	return &models.LinkData{
		Short:  suffix,
		Source: link,
	}, nil
}

func generateSuffix(link string) string {
	hash := sha1.Sum([]byte(link))
	return base64.StdEncoding.EncodeToString(hash[:8])
}
