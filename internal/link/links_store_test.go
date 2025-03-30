package link_test

import (
	"errors"
	"sync"

	"github.com/user-xat/short-link/internal/models"
)

type LinksStoreMap struct {
	sync.RWMutex
	store map[uint]models.Link
	id    uint
}

// Create new links map store
func NewLinksStoreMap() *LinksStoreMap {
	return &LinksStoreMap{
		store: make(map[uint]models.Link),
		id:    0,
	}
}

// Add new record in map store
func (ls *LinksStoreMap) Create(link *models.Link) (*models.Link, error) {
	ls.Lock()
	defer ls.Unlock()

	link.ID = ls.id
	ls.id++
	ls.store[link.ID] = *link
	return link, nil
}

// Get record from map store by key
func (ls *LinksStoreMap) GetByHash(hash string) (*models.Link, error) {
	ls.RLock()
	defer ls.RUnlock()

	for _, value := range ls.store {
		if value.Hash == hash {
			return &value, nil
		}
	}
	return nil, errors.New("not found")
}

func (ls *LinksStoreMap) GetById(id uint) (*models.Link, error) {
	return &models.Link{}, nil
}

func (ls *LinksStoreMap) GetAll(limit, offset int) []models.Link {
	return make([]models.Link, 0)
}

func (ls *LinksStoreMap) Update(link *models.Link) (*models.Link, error) {
	return link, nil
}

func (ls *LinksStoreMap) Delete(id uint) error {
	return nil
}

func (ls *LinksStoreMap) Count() int64 {
	return int64(len(ls.store))
}
