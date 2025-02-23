package main

import (
	"context"
	"crypto/sha1"
	"encoding/base64"

	"github.com/user-xat/short-link/pkg/models"
)

// Interface for working with the link storage
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

// Return data about link
func (sl *ShortLink) Get(ctx context.Context, suffix string) (*models.LinkData, error) {
	return sl.store.Get(ctx, suffix)
}

// Generates a short link, save it and returns the object LinkData
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

// Generates a hash based on the sha1 function and returns the first 8 bytes in the 64-base system as string.
func generateSuffix(link string) string {
	hash := sha1.Sum([]byte(link))
	return base64.StdEncoding.EncodeToString(hash[:8])
}
