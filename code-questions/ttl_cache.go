package main

import (
	"fmt"
	"sync"
	"time"
)

type item[V any] struct {
	value  V
	expiry time.Time
}

func (i item[V]) isExpiry() bool {
	return time.Now().After(i.expiry)
}

type TTLCache[K comparable, V any] struct {
	items map[K]item[V]
	mu    sync.Mutex
}

func NewTTLCache[K comparable, V any]() *TTLCache[K, V] {
	c := &TTLCache[K, V]{
		items: make(map[K]item[V]),
	}

	go func() {
		for range time.Tick(5 * time.Second) {
			c.mu.Lock()
			for key, item := range c.items {
				if item.isExpiry() {
					delete(c.items, key)
				}
			}
			c.mu.Unlock()
		}
	}()

	return c
}

func (c *TTLCache[K, V]) Set(key K, value V, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = item[V]{
		value:  value,
		expiry: time.Now().Add(ttl),
	}
}

func (c *TTLCache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, found := c.items[key]
	if !found {
		return item.value, found
	}
	if item.isExpiry() {
		delete(c.items, key)
		return item.value, false
	}
	return item.value, found
}

func (c *TTLCache[K, V]) Remove(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}

func (c *TTLCache[K, V]) Pop(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, found := c.items[key]
	if !found {
		return item.value, found
	}

	if item.isExpiry() {

		return item.value, false
	}

	delete(c.items, key)

	return item.value, found
}

func main() {
	// Create a new TTLCache instance
	myTTLCache := NewTTLCache[string, int]()

	// Set key-value pairs with TTL in the cache
	myTTLCache.Set("one", 1, 5*time.Second)
	myTTLCache.Set("two", 2, 10*time.Second)
	myTTLCache.Set("three", 3, 15*time.Second)

	// Retrieve values from the cache
	value, found := myTTLCache.Get("two")
	if found {
		fmt.Printf("Value for key 'two': %v\n", value)
	} else {
		fmt.Println("Key 'two' not found in the cache or has expired")
	}

	// Wait for a while to allow some items to expire
	time.Sleep(7 * time.Second)

	// Try to retrieve an expired key
	expiredValue, found := myTTLCache.Get("one")
	if found {
		fmt.Printf("Value for key 'one': %v\n", expiredValue)
	} else {
		fmt.Println("Key 'one' not found in the cache or has expired")
	}

	// Pop a key from the cache
	poppedValue, found := myTTLCache.Pop("two")
	if found {
		fmt.Printf("Popped value for key 'two': %v\n", poppedValue)
	} else {
		fmt.Println("Key 'two' not found in the cache or has expired")
	}

	// Remove a key from the cache
	myTTLCache.Remove("three")
}
