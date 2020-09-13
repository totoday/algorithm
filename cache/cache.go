package cache

type node struct {
	Key string
	Val  interface{}
	Prev *node
	Next *node
}

type Cache interface {
	Put(key string, val interface{})
	Get(key string) interface{}
	Remove(key string)
}
