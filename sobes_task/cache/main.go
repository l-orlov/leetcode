/*
   Simple cache

   Constructor

   Get
   Set
   Delete

   поле TTL - можно использовать для времени жизни ключа
   Cache Garbage collector
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	delDelay        = 5 * time.Second
	elementLifeTime = 10 * time.Second
)

type (
	CacheElement struct {
		Val       string
		CreatedAt time.Time
	}
	Cache struct {
		data            map[string]*CacheElement
		mu              *sync.RWMutex
		delDelay        time.Duration
		elementLifeTime time.Duration
	}
)

func NewCache() *Cache {
	data := make(map[string]*CacheElement)
	cache := Cache{
		data:            data,
		mu:              &sync.RWMutex{},
		delDelay:        delDelay,
		elementLifeTime: elementLifeTime,
	}

	return &cache
}

func (c *Cache) Get(key string) *CacheElement {
	c.mu.RLock()
	defer c.mu.RUnlock()

	el, ok := c.data[key]
	if !ok {
		return nil
	}

	return el
}

func (c *Cache) Set(key string, value *CacheElement) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
}

func (c *Cache) DeleteByKeys(keys []string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, k := range keys {
		delete(c.data, k)
	}
}

func (c *Cache) startClear() {
	ticker := time.NewTicker(c.delDelay)

	for {
		select {
		case <-ticker.C:
			c.mu.RLock()

			var keysToDelete []string
			for k, v := range c.data {
				if time.Since(v.CreatedAt) > c.elementLifeTime {
					keysToDelete = append(keysToDelete, k)
				}
			}

			c.mu.RUnlock()
			c.DeleteByKeys(keysToDelete)
		}
	}
}

func main() {
	fmt.Println("Hello, World!")
}
