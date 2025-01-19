package memcached

import (
	"context"
	"errors"
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/user-xat/short-link-server/pkg/models"
)

// Wrapper for Memcached
type CachedLinkModel struct {
	mc *memcache.Client
}

func NewCachedLinkModel(server ...string) *CachedLinkModel {
	return &CachedLinkModel{
		mc: memcache.New(server...),
	}
}

// Insert data into cache by short link
func (m *CachedLinkModel) Set(ctx context.Context, link *models.LinkData) (string, error) {
	res := make(chan error)
	go func(res chan<- error) {
		err := m.mc.Set(&memcache.Item{
			Key:        link.Short,
			Value:      []byte(link.Source),
			Expiration: 300,
		})

		res <- err
	}(res)

	var err error
	select {
	case <-ctx.Done():
		return "", errors.New("failed to add data in memcached")
	case e := <-res:
		err = e
	}

	if err != nil {
		return "", err
	}

	return link.Short, nil
}

// Get data from cache by key
func (m *CachedLinkModel) Get(ctx context.Context, key string) (*models.LinkData, error) {
	type resData struct {
		item *memcache.Item
		err  error
	}
	res := make(chan resData)

	go func(res chan<- resData) {
		data, err := m.mc.Get(key)
		res <- resData{
			item: data,
			err:  err,
		}
	}(res)

	var err error
	var item *memcache.Item
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("failed to get data from memcached: %v", ctx.Err())
	case r := <-res:
		item, err = r.item, r.err
	}

	if err != nil {
		return nil, err
	}

	return &models.LinkData{
		Short:  item.Key,
		Source: string(item.Value),
	}, nil
}

// Delete all data from cache
func (m *CachedLinkModel) DeleteAll(ctx context.Context) error {
	res := make(chan error)
	go func(r chan<- error) {
		r <- m.mc.DeleteAll()
	}(res)

	select {
	case <-ctx.Done():
		return fmt.Errorf("failed to delete all data from memcached: %v", ctx.Err())
	case err := <-res:
		return err
	}
}
