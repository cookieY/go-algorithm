package cache

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "k3"
	v1, v2, v3 := "value1", "value2", "v3"
	cap := len(k1 + k2 + v1 + v2 + k3 + v3)
	lruK := NewLruK(cap, nil)
	lruK.Set(k1, String(v1))
	lruK.Set(k2, String(v2))
	lruK.Get(k1)
	time.Sleep(time.Second * 5)
	lruK.Get(k1)

}
