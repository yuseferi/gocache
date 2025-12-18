package gocache

import (
	"testing"
	"time"
)

func TestCache_Get(t *testing.T) {
	cache := NewCache[string](time.Second)
	cache.Set("key", "value", time.Second)

	value, found := cache.Get("key")
	if !found {
		t.Errorf("Expected key 'key' to be found")
	}
	if value != "value" {
		t.Errorf("Expected value 'value', got %v", value)
	}
	// Test if it really expired
	time.Sleep(time.Second * 2)
	value, found = cache.Get("key")
	if found || value != "" {
		t.Errorf("Expected key 'key' to be expired and not found")
	}

	// Test a non-existent key
	value, found = cache.Get("nonexistent")
	if found || value != "" {
		t.Errorf("Expected non-existent key to not be found")
	}
}

func TestCache_Get_ExpiredItemBeforeCleanup(t *testing.T) {
	// Use a long cleanup period to ensure the cleanup goroutine doesn't run during the test
	cache := NewCache[string](time.Hour)

	// Set an item with a very short expiration
	cache.Set("key", "value", time.Millisecond)

	// Wait for the item to expire (but cleanup goroutine won't run)
	time.Sleep(time.Millisecond * 10)

	// The item should still exist in the cache map but be expired
	// Get should return zero value and false for expired items
	value, found := cache.Get("key")
	if found {
		t.Errorf("Expected expired key 'key' to not be found")
	}
	if value != "" {
		t.Errorf("Expected zero value for expired key, got %v", value)
	}
}

func TestCache_Set(t *testing.T) {
	cache := NewCache[string](time.Second)

	// Test setting a key-value pair
	cache.Set("key", "value", time.Minute)

	// Retrieve the value to verify
	value, found := cache.Get("key")
	if !found {
		t.Errorf("Expected key 'key' to be found")
	}
	if value != "value" {
		t.Errorf("Expected value 'value', got %v", value)
	}

	// Test updating an existing key
	cache.Set("key", "newvalue", time.Minute)

	// Retrieve the value to verify the update
	value, found = cache.Get("key")
	if !found {
		t.Errorf("Expected key 'key' to be found")
	}
	if value != "newvalue" {
		t.Errorf("Expected value 'newvalue', got %v", value)
	}
}

func TestCache_Delete(t *testing.T) {
	cache := NewCache[string](time.Second)
	cache.Set("key", "value", time.Minute)

	// Delete an existing key
	cache.Delete("key")

	// Verify the key is no longer found
	_, found := cache.Get("key")
	if found {
		t.Errorf("Expected key 'key' to be deleted")
	}

	// Delete a non-existent key
	cache.Delete("nonexistent")
}

func TestCache_Clear(t *testing.T) {
	cache := NewCache[string](time.Second)
	cache.Set("key1", "value1", time.Minute)
	cache.Set("key2", "value2", time.Minute)

	// Clear the cache
	cache.Clear()

	// Verify the cache is empty
	if size := cache.Size(); size != 0 {
		t.Errorf("Expected cache size 0, got %d", size)
	}
}

func TestCache_Size(t *testing.T) {
	cache := NewCache[string](time.Second)

	// Empty cache
	if size := cache.Size(); size != 0 {
		t.Errorf("Expected cache size 0, got %d", size)
	}

	// Non-empty cache
	cache.Set("key1", "value1", time.Minute)
	cache.Set("key2", "value2", time.Minute)
	if size := cache.Size(); size != 2 {
		t.Errorf("Expected cache size 2, got %d", size)
	}
}
