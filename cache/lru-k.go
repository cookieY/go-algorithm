package cache

import (
	"github.com/cookieY/go-algorithm/cache/simpleLru"
	"sync"
)

type LruK struct {
	first simpleLru.LRU
	last  simpleLru.LRU
	lock  sync.RWMutex
}

func NewLruK(maxBytes int, evicted func(key string, value simpleLru.CacheValue)) *LruK {
	return &LruK{
		first: simpleLru.New(simpleLru.LruConfig{Evicted: nil, MaxBytes: maxBytes, IsHistory: true, WaitTime: 2}),
		last:  simpleLru.New(simpleLru.LruConfig{Evicted: evicted, MaxBytes: maxBytes}),
	}
}

func (l *LruK) Set(key string, value simpleLru.CacheValue) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.last.Contains(key) {
		l.last.Set(key, value)
		return
	}
	l.first.Set(key, value)

}

func (l *LruK) Get(key string) (value simpleLru.CacheValue, ok bool) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if value, ok := l.first.Get(key); ok {
		if l.first.GetCursor(key) >= 2 {
			l.last.Set(key, value)
			l.first.Delete(key)
		}
		return value, ok
	}

	if value, ok := l.last.Get(key); ok {
		return value, ok
	}
	return
}

func (l *LruK) Purge() {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.first.Purge()
	l.last.Purge()
}

func (l *LruK) Contains(key string) (ok bool) {
	l.lock.RLock()
	defer l.lock.RUnlock()
	if ok = l.first.Contains(key); ok {
		return
	}
	if ok = l.last.Contains(key); ok {
		return
	}
	return
}

func (l *LruK) Len() int {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.last.Len() + l.first.Len()
}

func (l *LruK) Delete(key string) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.first.Delete(key)
	l.last.Delete(key)
}
