package main

import "sync"

type cache struct {
	content map[string]string
	mu      sync.Mutex
}

func newCache() *cache {
	return &cache{
		content: make(map[string]string),
	}
}
