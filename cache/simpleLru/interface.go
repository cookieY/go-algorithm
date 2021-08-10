package simpleLru

type LRU interface {
	Get(key string) (value CacheValue, ok bool)
	Set(key string, value CacheValue)
	Purge()
	Contains(key string) (ok bool)
	Len() int
	Delete(key string)
	GetCursor(key string) (index int)
}
