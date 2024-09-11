package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		entries: map[string]cacheEntry{},
	}
	ticker := time.NewTicker(interval)
	go c.reapLoop(ticker)

	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.entries[key]
	if ok {
		return val.val, true
	} else {
		return val.val, false
	}
}

func (c *Cache) reapLoop(ticker *time.Ticker) {
	limit := time.Now()
	for {
		nextLimit := <-ticker.C
		c.mu.Lock()
		for key, entry := range c.entries {
			if entry.createdAt.Before(limit) {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
		limit = nextLimit
	}
}
