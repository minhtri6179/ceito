package caching

import (
	"time"

	"github.com/allegro/bigcache/v3"
)

type Caching struct {
	cache *bigcache.BigCache
}

func NewCaching(t time.Duration) *Caching {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(t))
	return &Caching{cache: cache}
}
