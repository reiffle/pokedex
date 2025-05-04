package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu    *sync.Mutex
	cache map[string]cacheEntry
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{val: val}
	c.cache[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cache[key]
	if ok {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval) //specify a ticker interval
	for range ticker.C {               //run this process every time the interval passes
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}
