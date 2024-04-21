package lru

import "container/list"

type Cache struct {
	maxBytes  int64 // 最大的存储量，单位byte（若为0，表示没有上限）
	nBytes    int64 // 当前使用的存储量，单位byte
	ll        *list.List
	m         map[string]*list.Element
	OnEvicted func(key string, value Value)
}

/*
// Element 表示双向链表ll的节点，同时也是字典m键值对中的值
// Element's Value 我们用自定义的entry
type Element struct {
	next, prev *Element

	// The list to which this element belongs.
	list *List

	// The value stored with this element.
	Value any
}
*/

type entry struct {
	key   string
	value Value
}

// Value 为了通用性，实现Len即可
// 语义是返回占用的内存
type Value interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		nBytes:    0,
		ll:        list.New(),
		m:         make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (value Value, ok bool) {
	if element, ok := c.m[key]; ok {
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
	delete(c.m, kv.key)
	c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}
}

func (c *Cache) Set(key string, value Value) {
	ele, ok := c.m[key]
	if ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		kv.value = value
		// 更新当前lru的总存储量
		c.nBytes = c.nBytes + int64(value.Len()) - int64(kv.value.Len())
	} else {
		ele = c.ll.PushFront(&entry{key, value})
		c.nBytes += int64(len(key)) + int64(value.Len())
		c.m[key] = ele
	}

	// 如果超出内存限制，进行元素驱逐
	for c.maxBytes != 0 && c.maxBytes < c.nBytes {
		c.RemoveOldest()
	}

	return
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
