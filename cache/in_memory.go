package cache

import (
	"context"
	"time"

	goCache "github.com/patrickmn/go-cache"
)

type MemoryCache struct {
	cache *goCache.Cache
	ttl   time.Duration
}

func InMemoryCache(ttl time.Duration) *MemoryCache {
	c := goCache.New(ttl, 2*ttl)

	return &MemoryCache{cache: c, ttl: ttl}
}

func (c *MemoryCache) Add(ctx context.Context, key string, value interface{}) {
	c.cache.Set(key, value, c.ttl)
}

func (c *MemoryCache) Get(ctx context.Context, key string) (interface{}, bool) {
	s, found := c.cache.Get(key)
	if !found {
		return struct{}{}, false
	}

	return s, true
}
