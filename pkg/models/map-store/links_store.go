package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/user-xat/short-link-server/pkg/models"
)

type LinksStoreMap struct {
	sync.RWMutex
	store map[string]string
}

func NewLinksStoreMap() *LinksStoreMap {
	return &LinksStoreMap{
		store: make(map[string]string),
	}
}

func (ls *LinksStoreMap) Set(ctx context.Context, link *models.LinkData) (string, error) {
	ls.Lock()
	defer ls.Unlock()

	ls.store[link.Short] = link.Source
	return link.Short, nil
}

func (ls *LinksStoreMap) Get(ctx context.Context, suffix string) (*models.LinkData, error) {
	ls.RLock()
	defer ls.RUnlock()

	link, ok := ls.store[suffix]
	if !ok {
		return nil, fmt.Errorf("link by suffix %s not found", suffix)
	}

	return &models.LinkData{
		Source: link,
		Short:  suffix,
	}, nil
}
