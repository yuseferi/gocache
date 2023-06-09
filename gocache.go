// Package gocache provides a data race-free cache implementation in Go.
//
// Usage:
//
//	cache := gocache.NewCache(time.Minute * 2) // with 2 minutes interval cleaning
//	cache.Set("key", "value", time.Minute)
//	value, found := cache.Get("key")
//	cache.Delete("key")
//	cache.Clear()
//	size := cache.Size()
package gocache

import (
	"sync"
	"time"
)

// Cache represents a data race-free cache.
type Cache struct {
	items                map[string]cacheItem
	mutex                sync.RWMutex
	cleanupExpiredPeriod time.Duration
}

type cacheItem struct {
	value      interface{}
	expiration time.Time
}

// NewCache creates a new Cache instance.
func NewCache(cleanupExpiredPeriod time.Duration) *Cache {
	cache := &Cache{
		items: make(map[string]cacheItem),
	}
	cache.cleanupExpiredPeriod = cleanupExpiredPeriod
	// Start a goroutine to periodically check for expired items and remove them
	go cache.deleteExpiredItems()

	return cache
}

// Get retrieves the value associated with the specified key from the cache.
// It returns the value and a boolean indicating whether the key was found or not.
// If the key is found but the associated item has expired, the value will be nil
// and the boolean will be false.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, false
	}

	if item.expiration.Before(time.Now()) {
		return nil, false
	}

	return item.value, true
}

// Set adds or updates a key-value pair in the cache with the specified expiration duration.
// If the key already exists, its value and expiration are updated.
func (c *Cache) Set(key string, value interface{}, expiration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	expirationTime := time.Now().Add(expiration)
	c.items[key] = cacheItem{
		value:      value,
		expiration: expirationTime,
	}
}

// Delete removes the specified key and its associated value from the cache.
// If the key does not exist in the cache, the function does nothing.
func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.items, key)
}

// Clear removes all items from the cache, making it empty.
func (c *Cache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.items = make(map[string]cacheItem)
}

// Size returns the number of items currently stored in the cache.
func (c *Cache) Size() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return len(c.items)
}

// deleteExpiredItems is a background goroutine that periodically checks for expired items in the cache
// and removes them. It runs indefinitely after the Cache is created.
func (c *Cache) deleteExpiredItems() {
	for {
		<-time.After(c.cleanupExpiredPeriod) // Adjust the time interval for checking expired items

		c.mutex.Lock()
		for key, item := range c.items {
			if item.expiration.Before(time.Now()) {
				delete(c.items, key)
			}
		}
		c.mutex.Unlock()
	}
}
