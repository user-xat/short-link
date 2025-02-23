package mapstore

import (
	"context"
	"sync"

	"github.com/user-xat/short-link/pkg/models"
)

type LinksStoreMap struct {
	sync.RWMutex
	store map[string]string
}

// Create new links map store
func NewLinksStoreMap() *LinksStoreMap {
	return &LinksStoreMap{
		store: make(map[string]string),
	}
}

// Add new record in map store
func (ls *LinksStoreMap) Set(ctx context.Context, link *models.LinkData) (string, error) {
	ls.Lock()
	defer ls.Unlock()

	ls.store[link.Short] = link.Source
	return link.Short, nil
}

// Get record from map store by key
func (ls *LinksStoreMap) Get(ctx context.Context, key string) (*models.LinkData, error) {
	ls.RLock()
	defer ls.RUnlock()

	link, ok := ls.store[key]
	if !ok {
		return nil, models.ErrNotRecord
	}

	return &models.LinkData{
		Source: link,
		Short:  key,
	}, nil
}
