package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var ErrNilCache = errors.New("key does not exist")

type cacheItem struct {
	value       any
	ttlDeadline int64
}

type cache struct {
	data map[string]cacheItem
	mu   sync.RWMutex
	ctx  context.Context
}

func NewCache(ctx context.Context, startSize int) *cache {
	c := &cache{
		data: make(map[string]cacheItem, startSize),
		ctx:  ctx,
	}
	go c.invalidationLoop()
	return c
}

func (c *cache) Add(key string, value any, ttl time.Duration) error {
	deadline := time.Now().Add(ttl).UnixNano()
	c.mu.Lock()
	c.data[key] = cacheItem{value: value, ttlDeadline: deadline}
	c.mu.Unlock()
	return nil
}

func (c *cache) Get(key string) (any, error) {
	c.mu.RLock()
	cacheItem, ok := c.data[key]
	c.mu.RUnlock()
	if !ok {
		return nil, ErrNilCache
	}
	return cacheItem.value, nil
}

func (c *cache) Delete(key string) {
	c.mu.Lock()
	delete(c.data, key)
	c.mu.Unlock()
}

func (c *cache) invalidationLoop() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-c.ctx.Done():
			return
		case t := <-ticker.C:
			fmt.Println("current cache", c.data)
			nowNano := t.UnixNano()
			c.mu.Lock()
			for key, value := range c.data {
				if value.ttlDeadline <= nowNano {
					delete(c.data, key)
				}
			}
			c.mu.Unlock()
		}
	}

	
}

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
// 	cacheStartSize := 512

// 	cache := NewCache(ctx, cacheStartSize)

// 	cache.Add("rus", "champ", 10*time.Second)
// 	cache.Add("bla", "foo", 3*time.Second)

// 	for {
// 		time.Sleep(1 * time.Second)
// 	}

// }
