package link_test

import (
	"errors"
	"sync"

	"github.com/user-xat/short-link/internal/link"
)

type LinksStoreMap struct {
	sync.RWMutex
	store map[uint]link.Link
	id    uint
}

// Create new links map store
func NewLinksStoreMap() *LinksStoreMap {
	return &LinksStoreMap{
		store: make(map[uint]link.Link),
		id:    0,
	}
}

// Add new record in map store
func (ls *LinksStoreMap) Create(link *link.Link) (*link.Link, error) {
	ls.Lock()
	defer ls.Unlock()

	link.ID = ls.id
	ls.id++
	ls.store[link.ID] = *link
	return link, nil
}

// Get record from map store by key
func (ls *LinksStoreMap) GetByHash(hash string) (*link.Link, error) {
	ls.RLock()
	defer ls.RUnlock()

	for _, value := range ls.store {
		if value.Hash == hash {
			return &value, nil
		}
	}
	return nil, errors.New("not found")
}

func (ls *LinksStoreMap) GetById(id uint) (*link.Link, error) {
	return &link.Link{}, nil
}

func (ls *LinksStoreMap) GetAll(limit, offset int) []link.Link {
	return make([]link.Link, 0)
}

func (ls *LinksStoreMap) Update(link *link.Link) (*link.Link, error) {
	return link, nil
}

func (ls *LinksStoreMap) Delete(id uint) error {
	return nil
}

func (ls *LinksStoreMap) Count() int64 {
	return int64(len(ls.store))
}
