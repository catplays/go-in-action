package geechache

import (
	"catwang.com/go-in-action/geechache/lru"
	"sync"
)

type cache struct {
	mux sync.Mutex
	lru *lru.Cache
	cacheBytes int64
}

func (c *cache) add(key string, value ByteView)  {
	c.mux.Lock()
	defer c.mux.Unlock()
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}
	c.lru.Add(key, value)

}

func (c *cache) get(key string) (value ByteView, b bool)  {
	c.mux.Lock()
	defer c.mux.Unlock()
	if c.lru == nil {
		return
	}
	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), true
	}
	return
}
