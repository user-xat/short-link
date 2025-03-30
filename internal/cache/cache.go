package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/user-xat/short-link/internal/models"
	"github.com/user-xat/short-link/pkg/event"
)

type CacheConfig struct {
	Addr        string        `yaml:"addr"`
	Password    string        `yaml:"password"`
	User        string        `yaml:"user"`
	DB          int           `yaml:"db"`
	MaxRetries  int           `yaml:"max_retries"`
	DialTimeout time.Duration `yaml:"dial_timeout"`
	Timeout     time.Duration `yaml:"timeout"`
}

type CacheDeps struct {
	Config   CacheConfig
	Ctx      context.Context
	TTL      time.Duration
	EventBus *event.EventBus
}

type Cache struct {
	db       *redis.Client
	ttl      time.Duration
	eventBus *event.EventBus
}

func NewCache(deps CacheDeps) (*Cache, error) {
	db := redis.NewClient(&redis.Options{
		Addr: deps.Config.Addr,
		// Password:     cfg.Password,
		// DB:           cfg.DB,
		// Username:     cfg.User,
		// MaxRetries:   cfg.MaxRetries,
		// DialTimeout:  cfg.DialTimeout,
		// ReadTimeout:  cfg.Timeout,
		// WriteTimeout: cfg.Timeout,
	})
	if err := db.Ping(deps.Ctx).Err(); err != nil {
		log.Printf("failed to connect to redis server: %s\n", err.Error())
		return nil, err
	}
	return &Cache{
		db:       db,
		ttl:      deps.TTL,
		eventBus: deps.EventBus,
	}, nil
}

// Get link data from redis store by key
func (c *Cache) Get(ctx context.Context, hash string) (*models.Link, error) {
	val, err := c.db.Get(ctx, hash).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, errors.New("not found")
	} else if err != nil {
		return nil, fmt.Errorf("redis error: %v", err)
	}
	var link models.Link
	err = json.Unmarshal(val, &link)
	if err != nil {
		return nil, err
	}
	return &link, nil
}

// Add record to redis store, returns key
func (c *Cache) Set(ctx context.Context, link *models.Link) error {
	jsonData, err := json.Marshal(link)
	if err != nil {
		return err
	}
	if err := c.db.Set(ctx, link.Hash, jsonData, c.ttl).Err(); err != nil {
		return fmt.Errorf("failed to insert data into redis: %v", err)
	}
	return nil
}

func (c *Cache) Delete(ctx context.Context, hash string) error {
	if err := c.db.Del(ctx, hash).Err(); err != nil {
		return err
	}
	return nil
}

func (c *Cache) UpdateCache(msg *event.Event) {
	switch msg.Type {
	case event.EventLinkUpdated:
		fallthrough
	case event.EventLinkVisited:
		link, ok := msg.Data.(*models.Link)
		if !ok {
			log.Fatalln("Bad EventLinkUpdated Data: ", msg.Data)
		}
		c.Set(context.Background(), link)
	case event.EventLinkDeleted:
		link, ok := msg.Data.(*models.Link)
		if !ok {
			log.Fatalln("Bad EventLinkDeleted Data: ", msg.Data)
		}
		c.Delete(context.Background(), link.Hash)
	}
}
