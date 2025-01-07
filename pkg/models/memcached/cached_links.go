package memcached

import (
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
func (m *CachedLinkModel) Insert(link *models.CachedLink) error {
	return m.mc.Set(&memcache.Item{
		Key:        link.Short,
		Value:      []byte(link.Source),
		Expiration: 300,
	})
}

// Get data from cache by short link
func (m *CachedLinkModel) Get(short string) (*models.CachedLink, error) {
	data, err := m.mc.Get(short)
	if err != nil {
		return nil, err
	}

	return &models.CachedLink{
		Short:  data.Key,
		Source: string(data.Value),
	}, nil
}

// Clear cache
func (m *CachedLinkModel) DeleteAll() error {
	return m.mc.DeleteAll()
}
