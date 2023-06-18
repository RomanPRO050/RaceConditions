package main

import (
	"fmt"
	"sync"
	"time"
	_ "time"
)

type Cache struct {
	storage map[string]int
	mu      sync.RWMutex
}

func (c *Cache) Increase(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] += value
}

func (c *Cache) Set(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] = value
}

func (c *Cache) Get(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.storage[key]
}

func (c *Cache) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.storage, key)
}

const (
	k1   = "key1"
	step = 7
)

var wg sync.WaitGroup

func main() {
	cache := Cache{storage: make(map[string]int)}
	for i := 0; i < 5; i++ {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				cache.Increase(k1, step)
				time.Sleep(time.Millisecond * 100)
			}()
		}
		for i := 0; i < 10; i++ {
			go func(i int) {
				defer wg.Done()
				cache.Set(k1, step*i)
				time.Sleep(time.Millisecond * 100)
			}(i)
		}
	}
	fmt.Println(cache.Get(k1))
}
