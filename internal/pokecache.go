package main

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	cached map[string]cacheEntry
	mu *sync.RWMutex
}

func NewCache(cacheDuration time.Duration) {
	
	cacheData := Cache{
		cached: make(map[string]cacheEntry),
		mu: &sync.RWMutex{},
	}

	ticker := time.NewTicker(cacheDuration)

	for range ticker.C {
		cacheData.reapLoop()
	}

}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cached[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {

	c.mu.RLock()
	defer c.mu.RUnlock()
	value, err := c.cached[key]

	if err == false {
		return []byte{}, false
	}

	return value.val, true

}

func (c *Cache) reapLoop() {

	c.mu.Lock()
	defer c.mu.Unlock()
	
	for key := range c.cached { 
		_, err := c.cached[key]

		if err == false {
			return
		}
		
		delete(c.cached, key)
	}

	return

}