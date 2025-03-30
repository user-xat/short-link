package cache_test

import (
	"context"
	"testing"
	"time"

	"github.com/user-xat/short-link/internal/cache"
	"github.com/user-xat/short-link/internal/models"
	"gorm.io/gorm"
)

func TestRWData(t *testing.T) {
	cfg := cache.CacheConfig{
		Addr:        "localhost:6379",
		Password:    "test1234",
		User:        "testuser",
		DB:          0,
		MaxRetries:  5,
		DialTimeout: 10 * time.Second,
		Timeout:     5 * time.Second,
	}

	cache, err := cache.NewCache(cache.CacheDeps{
		Config: cfg,
		Ctx:    context.Background(),
		TTL:    time.Minute,
	})
	if err != nil {
		t.Skipf("can't create redis store: %v", err)
	}

	link := models.Link{
		Model: gorm.Model{ID: 1},
		Url:   "http://github.com/",
		Hash:  "Zchs_he4_3hg",
	}

	// insert data
	if err := cache.Set(context.Background(), &link); err != nil {
		t.Errorf("failed to set data, error: %s\n", err.Error())
	}

	// read data
	val, err := cache.Get(context.Background(), link.Hash)
	if err != nil {
		t.Errorf("failed to get value, error: %v\n", err.Error())
	}
	if val.Url != link.Url {
		t.Errorf("got value=%s, want test value", val.Url)
	}
}
