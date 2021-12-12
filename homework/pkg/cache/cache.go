package cache

import (
	"sync"
)

type Cache struct {
	// cache name
	name  string
	data  map[string]string
	_lock sync.RWMutex
}

func (c *Cache) Get(k string) (string, bool) {
	value, is_ok := c.data[k]
	if is_ok {
		return value, true
	} else {
		return "", false
	}
}

func (c *Cache) Set(k string, v string) {
	c._lock.RLock()
	defer c._lock.RUnlock()
	c.data[k] = v
}

func (c *Cache) Del(k string) {
	c._lock.RLock()
	defer c._lock.RUnlock()
	delete(c.data, k)
}

func New(cacheName string) *Cache {
	return &Cache{
		name: cacheName,
		data: make(map[string]string),
	}
}
