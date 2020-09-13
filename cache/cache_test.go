package cache

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	cache := NewLRUCache(3)
	fmt.Println(cache.Get("a"))
	fmt.Println(cache)
	cache.Put("a", 1)
	cache.Put("b", 2)
	fmt.Println(cache.Get("a"))
	cache.Put("c", 3)
	cache.Put("d", 4)
	fmt.Println(cache.Get("a"), cache.Get("b"))
}