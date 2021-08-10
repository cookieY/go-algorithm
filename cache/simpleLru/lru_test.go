package simpleLru

import (
	"testing"
)

type String string

func (s String) Len() int {
	return len(s)
}

func TestLruNew(t *testing.T) {

	k1, k2, k3 := "key1", "key2", "k3"
	v1, v2, v3 := "value1", "value2", "v3"
	cap := len(k1 + k2 + v1 + v2)
	lru := New(LruConfig{
		Evicted:   nil,
		MaxBytes:  cap,
		IsHistory: false,
	})
	lru.Set(k1, String(v1))
	lru.Set(k2, String(v2))
	if _, ok := lru.Get("key1"); ok {
	}
	lru.Set(k3, String(v3))

	//
	//if _, ok := lru.Get("key1"); ok || lru.Len() != twoQ {
	//	t.Fatalf("Removeoldest key1 failed")
	//}
}
