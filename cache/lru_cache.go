package cache

type lruCache struct {
	capacity int
	head     *node
	tail     *node
	index    map[string]*node
}

func NewLRUCache(capacity int) Cache {
	head := &node{}
	tail := &node{}
	head.Next = tail
	tail.Prev = head
	return &lruCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		index:    make(map[string]*node, 0),
	}
}

func (cache *lruCache) Put(key string, val interface{}) {
	n := cache.index[key]
	if n != nil {
		n.Val = val
		cache.removeNode(n)
		cache.putHead(n)
		return
	}

	if len(cache.index) == cache.capacity {
		rmn := cache.tail.Prev
		cache.removeNode(rmn)
		delete(cache.index, rmn.Key)
	}

	n = &node{Key: key, Val: val}
	cache.putHead(n)
	cache.index[key] = n
}

func (cache *lruCache) Get(key string) interface{} {
	node := cache.index[key]
	if node == nil {
		return nil
	}

	cache.removeNode(node)
	cache.putHead(node)

	return node.Val
}

func (cache *lruCache) Remove(key string) {
	node := cache.index[key]
	if node == nil {
		return
	}

	cache.removeNode(node)
	delete(cache.index, key)
}

func (cache *lruCache) removeNode(node *node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node.Prev = nil
	node.Next = nil
}

func (cache *lruCache) putHead(node *node) {
	node.Next = cache.head.Next
	node.Prev = cache.head
	node.Next.Prev = node
	node.Prev.Next = node
}

func (cache *lruCache) String() string {
	res := ""

	node := cache.head.Next
	res += "lru ->: "
	for node != cache.tail {
		res += node.Key + " "
		node = node.Next
	}
	res += "\n"
	node = cache.tail.Prev
	res += "lru <-: "
	for node != cache.head {
		res += node.Key + " "
		node = node.Prev
	}
	res += "\n"

	return res
}
