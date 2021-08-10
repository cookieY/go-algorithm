package cache

type String string

func (t String) Len() int {
	return len(t)
}
