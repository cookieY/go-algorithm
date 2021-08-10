package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLru(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "k3"
	v1, v2, v3 := "value1", "value2", "v3"
	cap := len(k1 + k2 + v1 + v2)
	lru := NewLru(cap, nil)
	lru.Set(k1, String(v1))
	lru.Set(k2, String(v2))
	_, ok := lru.Get(k1)
	assert := assert.New(t)
	assert.Equal(ok, true)
	lru.Get(k1)
	lru.Set(k3, String(v3))
	_, ok = lru.Get(k2)
	assert.Equal(ok, false)

}
