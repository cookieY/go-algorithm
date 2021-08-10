package simpleLru

import (
	"container/list"
	"fmt"
	"time"
)

type CacheValue interface {
	Len() int
}

type cursor struct {
	index   int
	preTime time.Time
}

type ele struct {
	key    string
	value  CacheValue
	option cursor
}

type simpleCache struct {
	maxBytes  int
	usedBytes int
	ll        *list.List
	cache     map[string]*list.Element
	evicted   func(key string, value CacheValue)
	isHistory bool
	isFIFO    bool
	waitTime  time.Duration
}

type LruConfig struct {
	Evicted   func(key string, value CacheValue)
	MaxBytes  int
	IsHistory bool
	IsFIFO    bool
	WaitTime  time.Duration
}

func New(config LruConfig) *simpleCache {
	return &simpleCache{
		maxBytes:  config.MaxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		evicted:   config.Evicted,
		isHistory: config.IsHistory,
		waitTime:  config.WaitTime,
	}
}

func (s *simpleCache) Get(key string) (value CacheValue, ok bool) {
	if ll, ok := s.cache[key]; ok {
		if !s.isFIFO {
			s.ll.MoveToFront(ll)
		}
		kv := s.cache[key].Value.(*ele)
		s.historyTableCheck(kv)
		return kv.value, true
	}
	return
}

func (s *simpleCache) GetCursor(key string) (index int) {
	if _, ok := s.cache[key]; ok {
		return s.cache[key].Value.(*ele).option.index
	}
	return 0
}

func (s *simpleCache) RemoveOldest() {
	if ll := s.ll.Back(); ll != nil {
		s.ll.Remove(ll)
		kv := ll.Value.(*ele)
		delete(s.cache, kv.key)
		fmt.Printf("remove key:%s\n", kv.key)
		if s.evicted != nil {
			s.evicted(kv.key, kv.value.(CacheValue))
		}
	}
}

func (s *simpleCache) Delete(key string) {
	if ll, ok := s.cache[key]; ok {
		s.ll.Remove(ll)
		delete(s.cache, key)
		if s.evicted != nil {
			s.evicted(key, ll.Value.(*ele).value)
		}
	}
}

func (s *simpleCache) Set(key string, value CacheValue) {
	if ll, ok := s.cache[key]; ok {
		if !s.isFIFO {
			s.ll.MoveToFront(ll)
		}
		kv := ll.Value.(*ele)
		s.usedBytes += value.Len() - kv.value.Len()
		kv.value = value
	} else {
		ll := s.ll.PushFront(&ele{
			key:   key,
			value: value,
			option: cursor{
				index:   0,
				preTime: time.Now(),
			},
		})
		s.cache[key] = ll
		s.usedBytes += len(key) + value.Len()
	}

	if s.maxBytes != 0 && s.maxBytes < s.usedBytes {
		s.RemoveOldest()
	}
}

func (s *simpleCache) Purge() {
	for k, v := range s.cache {
		if s.evicted != nil {
			s.evicted(k, v.Value.(*ele).value)
		}
		delete(s.cache, k)
	}
	s.ll.Init()
}

func (s *simpleCache) Contains(key string) (ok bool) {
	_, ok = s.cache[key]
	return ok
}

func (s *simpleCache) Len() int {
	return s.usedBytes
}

func (s *simpleCache) historyTableCheck(kv *ele) {
	if s.isHistory {
		if time.Now().Before(kv.option.preTime.Add(time.Second * s.waitTime)) {
			kv.option.index++
			return
		}
		kv.option.index = 1
	}
}
