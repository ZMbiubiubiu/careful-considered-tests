package main

import (
	"math/rand"
)

type Skiplist struct {
	head *node
}

type node struct {
	key, val int
	nexts    []*node
}

// Get 根据key寻找对应的val，第二个返回值代表是否存在key
func (s *Skiplist) Get(key int) (int, bool) {
	if _node := s.search(key); _node != nil {
		return _node.val, true
	}
	return -1, false
}

func (s *Skiplist) search(key int) *node {
	// 每次检索从头部出发
	move := s.head
	levels := len(s.head.nexts)
	for level := levels - 1; level >= 0; level-- {
		// 一直向右遍历
		// tips：这里可能有点不好理解，move.nexts[level] == 本层的move.next
		for move.nexts[level] != nil && move.nexts[level].key < key {
			move = move.nexts[level]
		}

		// 找到了key
		if move.nexts[level] != nil && move.nexts[level].key == key {
			return move.nexts[level]
		}
	}
	return nil
}

func (s *Skiplist) roll() int {
	var level int
	for rand.Int() > 0 {
		level++
	}
	return level
}

func (s *Skiplist) Put(key, val int) {
	// if key,val already exist, we can update directly
	if _node := s.search(key); _node != nil {
		_node.val = val
		return
	}

	// roll新高度
	level := s.roll()

	for len(s.head.nexts)-1 < level {
		s.head.nexts = append(s.head.nexts, nil)
	}

	//	创建出的新节点
	newNode := node{
		key:   key,
		val:   val,
		nexts: make([]*node, level+1),
	}

	move := s.head
	for level := level; level >= 0; level-- {
		// 向右遍历，直到右侧节点不存在或者key值大于key
		for move.nexts[level] != nil && move.nexts[level].key < key {
			move = move.nexts[level]
		}

		// 调整指针关系，完成新节点的插入
		newNode.nexts[level] = move.nexts[level]
		move.nexts[level] = &newNode
	}
}

func (s *Skiplist) Del(key int) {
	if _node := s.search(key); _node == nil {
		return
	}

	move := s.head
	for level := len(s.head.nexts) - 1; level >= 0; level-- {
		for move.nexts[level] != nil && move.nexts[level].key < key {
			move = move.nexts[level]
		}
	}
}

func main() {

}
