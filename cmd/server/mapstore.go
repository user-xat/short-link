package main

import (
	"crypto/md5"
	"encoding/hex"
	"sync"
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

func (ls *LinksStoreMap) Set(link string) string {
	ls.Lock()
	defer ls.Unlock()

	suffix := generateSuffix(link)
	ls.store[suffix] = link
	return suffix
}

func (ls *LinksStoreMap) Get(suffix string) (string, bool) {
	ls.RLock()
	defer ls.RUnlock()

	link, ok := ls.store[suffix]
	return link, ok
}

func generateSuffix(link string) string {
	hash := md5.Sum([]byte(link))
	return hex.EncodeToString(hash[:])
}
