package internal

import (
	"sync"
	"time"
)

type CacheItem[T any] struct {
	Value     T
	ExpiresAt time.Time
}

type Cache[T any] struct {
	mu    sync.RWMutex
	items map[string]CacheItem[T]
	ttl   time.Duration
}

func NewCache[T any](ttl time.Duration) *Cache[T] {
	return &Cache[T]{
		items: make(map[string]CacheItem[T]), // map[keyType]ValueType
		ttl:   ttl,
	}
}

func (c *Cache[T]) Get(key string) (T, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	if !found || time.Now().After(item.ExpiresAt) {
		var zero T
		return zero, false
	}
	return item.Value, true
}

func (c *Cache[T]) Set(key string, value T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item := CacheItem[T]{
		Value:     value,
		ExpiresAt: time.Now().Add(c.ttl),
	}

	c.items[key] = item
}

func (c *Cache[T]) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k := range c.items {
		delete(c.items, k)
	}
}

func (c *Cache[T]) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}

func (c *Cache[T]) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.items)
}

func (c *Cache[T]) StartCleanUp(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			now := time.Now()
			c.mu.Lock()
			for k, v := range c.items {
				if now.After(v.ExpiresAt) {
					delete(c.items, k)
				}
			}
			c.mu.Unlock()
		}
	}()
}
