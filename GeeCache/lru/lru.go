package lru

import "container/list"

type entry struct {
	key   string
	value Value
}

// Value 为了键值对，值的通用性，实现Len即可，返回占用的内存
type Value interface {
	Len() int
}

type Cache struct {
	maxBytes  int64 // 最大的存储量（若为0，表示没有上限）
	nBytes    int64 // 当前使用的存储量
	ll        *list.List
	cache     map[string]*list.Element
	OnEvicted func(key string, value Value)
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		nBytes:    0,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (value Value, ok bool) {
	if element, ok := c.cache[key]; ok {
		c.ll.MoveToFront(element)
		kv := element.Value.(*entry)
		value = kv.value
		return kv.value, true
	}
	return
}

func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele == nil {
		return
	}

	c.ll.Remove(ele)
	kv := ele.Value.(*entry)
	delete(c.cache, kv.key)
	c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}
}

func (c *Cache) Set(key string, value Value) {
	ele, ok := c.cache[key]
	if ok {
		kv := ele.Value.(*entry)
		c.ll.MoveToFront(ele)
		c.nBytes = c.nBytes + int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele = c.ll.PushFront(&entry{key, value})
		c.nBytes += int64(len(key)) + int64(value.Len())
		c.cache[key] = ele
	}

	for c.maxBytes != 0 && c.maxBytes < c.nBytes {
		c.RemoveOldest()
	}

	return
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
