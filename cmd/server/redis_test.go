package main

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestRWData(t *testing.T) {
	cfg := Config{
		Addr:        "localhost:6379",
		Password:    "test1234",
		User:        "testuser",
		DB:          0,
		MaxRetries:  5,
		DialTimeout: 10 * time.Second,
		Timeout:     5 * time.Second,
	}

	db, err := NewClient(context.Background(), cfg)
	if err != nil {
		t.Fatal(err)
	}

	// insert data
	if err := db.Set(context.Background(), "key", "test value", 0).Err(); err != nil {
		t.Errorf("failed to set data, error: %s\n", err.Error())
	}

	if err := db.Set(context.Background(), "key2", 3333, 30*time.Second).Err(); err != nil {
		t.Errorf("failed to set data, error: %s\n", err.Error())
	}

	// read data
	val, err := db.Get(context.Background(), "key").Result()
	if err == redis.Nil {
		t.Error("value not found")
	} else if err != nil {
		t.Errorf("failed to get value, error: %v\n", err.Error())
	}
	if val != "test value" {
		t.Errorf("got value=%s, want test value", val)
	}

	val2, err := db.Get(context.Background(), "key2").Result()
	if err == redis.Nil {
		t.Error("value not found")
	} else if err != nil {
		t.Errorf("failed to get value, error: %v\n", err.Error())
	}
	if val2 != "3333" {
		t.Errorf("got value=%s, want test value", val)
	}
}
