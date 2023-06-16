package main

import (
	"fmt"
	"runtime"
	"sync"
	_ "time"
)

type Cache struct {
	storage map[string]int
	mu      sync.Mutex
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
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.storage[key]
}

func (c *Cache) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.storage, key)
}

func main() {
	runtime.GOMAXPROCS(20)
	strg := &Cache{
		storage: map[string]int{"Феррари": 240, "Мерседес": 120, "Ауди": 300},
	}
	go strg.Increase("Феррари", 50)
	go strg.Set("Ауди", 340)
	go strg.Set("Ауди", 350)
	go strg.Remove("Мерседес")
	for k, v := range strg.storage {
		fmt.Printf(k, v)
	}
}
