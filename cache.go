package main

import (
	"fmt"
	"sync"
)

type cache struct {
	content map[string]string
	mu      sync.Mutex
}

func newCache() *cache {
	return &cache{
		content: make(map[string]string),
	}
}

func (c *cache) entryExists(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, exists := c.content[key]

	return exists
}

func (c *cache) addEntry(key, value string) error {
	if c.entryExists(key) {
		return fmt.Errorf("key: %s exists in cache", key)
	}

	c.mu.Lock()
	c.content[key] = value
	c.mu.Unlock()

	return nil
}

func (c *cache) getEntry(key string) (string, error) {
	c.mu.Lock()
	res, exists := c.content[key]
	c.mu.Unlock()

	if !exists {
		return "", fmt.Errorf("no entry for key: %s", key)
	}

	return res, nil
}
