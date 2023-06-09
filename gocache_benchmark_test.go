package gocache

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func BenchmarkCache_Set(b *testing.B) {
	cache := NewCache(time.Minute)
	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		value := fmt.Sprintf("value%d", i)
		cache.Set(key, value, time.Minute)
	}
}

func BenchmarkCache_Get(b *testing.B) {
	cache := NewCache(time.Minute)
	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		value := fmt.Sprintf("value%d", i)
		cache.Set(key, value, time.Minute)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		cache.Get(key)
	}
}

func BenchmarkCache_Delete(b *testing.B) {
	cache := NewCache(time.Minute)
	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		value := fmt.Sprintf("value%d", i)
		cache.Set(key, value, time.Minute)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		cache.Delete(key)
	}
}

func BenchmarkCache_ConcurrentAccess(b *testing.B) {
	cache := NewCache(time.Minute)
	concurrency := 100
	numOperations := b.N

	// Populate the cache with initial data
	for i := 0; i < concurrency; i++ {
		key := strconv.Itoa(i)
		value := fmt.Sprintf("value%d", i)
		cache.Set(key, value, time.Minute)
	}

	// Run concurrent access to the cache
	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < numOperations; j++ {
				key := strconv.Itoa(j % concurrency)
				cache.Get(key)
			}
		}()
	}

	wg.Wait()
}
