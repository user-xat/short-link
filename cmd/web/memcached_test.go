package main

import (
	"testing"

	"github.com/bradfitz/gomemcache/memcache"
)

func TestRWData(t *testing.T) {
	mc := memcache.New("localhost:11211")
	defer mc.Close()

	// insert data
	if err := mc.Set(&memcache.Item{Key: "key", Value: []byte("test value")}); err != nil {
		t.Errorf("failed to set data, error: %s\n", err.Error())
	}

	if err := mc.Set(&memcache.Item{Key: "key2", Value: []byte("3333"), Expiration: 30}); err != nil {
		t.Errorf("failed to set data, error: %s\n", err.Error())
	}

	// read data
	val, err := mc.Get("key")
	if err != nil {
		t.Errorf("failed to get value, error: %v\n", err.Error())
	}
	if string(val.Value) != "test value" {
		t.Errorf("got value=%s, want test value", val.Value)
	}

	val2, err := mc.Get("key2")
	if err != nil {
		t.Errorf("failed to get value, error: %v\n", err.Error())
	}
	if string(val2.Value) != "3333" {
		t.Errorf("got value=%s, want test value", val.Value)
	}
}
