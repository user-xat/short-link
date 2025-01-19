package redis

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/user-xat/short-link-server/pkg/models"
)

type Config struct {
	Addr        string        `yaml:"addr"`
	Password    string        `yaml:"password"`
	User        string        `yaml:"user"`
	DB          int           `yaml:"db"`
	MaxRetries  int           `yaml:"max_retries"`
	DialTimeout time.Duration `yaml:"dial_timeout"`
	Timeout     time.Duration `yaml:"timeout"`
}

type LinkStoreRedis struct {
	db *redis.Client
}

func NewLinkStoreRedis(ctx context.Context, cfg Config) (*LinkStoreRedis, error) {
	db := redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
		// Password:     cfg.Password,
		// DB:           cfg.DB,
		// Username:     cfg.User,
		// MaxRetries:   cfg.MaxRetries,
		// DialTimeout:  cfg.DialTimeout,
		// ReadTimeout:  cfg.Timeout,
		// WriteTimeout: cfg.Timeout,
	})

	if err := db.Ping(ctx).Err(); err != nil {
		log.Printf("failed to connect to redis server: %s\n", err.Error())
		return nil, err
	}

	return &LinkStoreRedis{
		db: db,
	}, nil
}

func (s *LinkStoreRedis) Get(ctx context.Context, short string) (*models.LinkData, error) {
	link, err := s.db.Get(ctx, short).Result()
	if errors.Is(err, redis.Nil) {
		return nil, models.ErrNotRecord
	} else if err != nil {
		return nil, fmt.Errorf("redis error: %v", err)
	}

	return &models.LinkData{
		Source: link,
		Short:  short,
	}, nil
}

func (s *LinkStoreRedis) Set(ctx context.Context, link *models.LinkData) (string, error) {
	if err := s.db.Set(ctx, link.Short, link.Source, 0).Err(); err != nil {
		return "", fmt.Errorf("failed to insert data into redis: %v", err)
	}

	return link.Short, nil
}
