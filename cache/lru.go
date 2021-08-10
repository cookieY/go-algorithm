package cache

import (
	"github.com/cookieY/go-algorithm/cache/simpleLru"
	"sync"
)

type Lru struct {
	ll   simpleLru.LRU
	lock sync.RWMutex
}

func NewLru(maxBytes int, evicted func(key string, value simpleLru.CacheValue)) *Lru {
	return &Lru{
		ll: simpleLru.New(simpleLru.LruConfig{Evicted: evicted, MaxBytes: maxBytes}),
	}
}

func (l *Lru) Set(key string, value simpleLru.CacheValue) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.ll.Set(key, value)
}

func (l *Lru) Get(key string) (value simpleLru.CacheValue, ok bool) {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.ll.Get(key)
}
