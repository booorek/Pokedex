package pokecache

import (
	"sync"
	"time"
)

func NewCache(duration time.Duration) *Cache {
	c := &Cache{
		entries:  make(map[string]cacheEntry),
		duration: duration,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	newCacheEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.entries[key] = newCacheEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	elem, err := c.entries[key]
	if err {
		return nil, false
	}
	return elem.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.duration * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		time := time.Now()
		c.mu.Lock()
		for key, elem := range c.entries {
			if time.Sub(elem.createdAt) > c.duration {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}

type Cache struct {
	mu       sync.Mutex
	entries  map[string]cacheEntry
	duration time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
