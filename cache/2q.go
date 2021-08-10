package cache

import (
	"github.com/cookieY/go-algorithm/cache/simpleLru"
	"sync"
)

type TwoQ struct {
	fifo simpleLru.LRU
	lru  simpleLru.LRU
	lock sync.RWMutex
}

func New2Q(maxBytes int, evicted func(key string, value simpleLru.CacheValue)) *TwoQ {
	return &TwoQ{
		fifo: simpleLru.New(simpleLru.LruConfig{
			Evicted:  nil,
			MaxBytes: maxBytes,
			IsFIFO:   true,
		}),
		lru: simpleLru.New(simpleLru.LruConfig{
			Evicted:  evicted,
			MaxBytes: maxBytes,
		}),
	}
}

func (s *TwoQ) Get(key string) (value simpleLru.CacheValue, ok bool) {
	if value, ok := s.fifo.Get(key); ok {
		s.fifo.Delete(key)
		s.lru.Set(key, value)
		return value, ok
	}
	if value, ok := s.lru.Get(key); ok {
		return value, ok
	}
	return
}

func (s *TwoQ) Set(key string, value simpleLru.CacheValue) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.lru.Contains(key) {
		s.lru.Set(key, value)
		return
	}
	s.fifo.Set(key, value)
}
