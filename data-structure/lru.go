package main

import (
	"container/list"
)

type LRU struct {
	innerList *list.List            // list
	innerMap  map[int]*list.Element // map，之所以使用list.Element是为了方便调整指针
	size      int                   // 链表的最大长度，超过需要丢弃老节点
}

type entry struct {
	key   int
	value int
}

func (lru *LRU) get(key int) (int, bool) {
	// 1.如果剪枝不存在，直接返回
	e, ok := lru.innerMap[key]
	// fast path 可以快速返回的路径
	if !ok {
		return -1, false
	}
	// 2.将找到的元素插入到链表头部
	lru.innerList.MoveToFront(e)
	return e.Value.(*entry).value, true
}

func (lru *LRU) put(key, value int) (evicted bool) {
	e, ok := lru.innerMap[key]
	if ok {
		lru.innerList.MoveToFront(e)
		lru.innerMap[key].Value.(*entry).value = value
		return false
	}

	// 插入map
	ent := &entry{key, value}
	e = lru.innerList.PushFront(ent)
	lru.innerMap[key] = e

	if len(lru.innerMap) > lru.size {
		last := lru.innerList.Back()
		lru.innerList.Remove(last)
		// 注意：是 last 节点的 key
		delete(lru.innerMap, last.Value.(*entry).key)
		return true
	}
	return false
}

func newLRU(size int) *LRU {
	return &LRU{
		size:      size,
		innerList: list.New(),                        // 提供O(1)的快速插入
		innerMap:  make(map[int]*list.Element, size), // 提供O(1)的快速查找
	}
}

func main() {
	lru := newLRU(2)

	for i := 0; i < 100; i++ {
		lru.put(i, i)
	}
}
