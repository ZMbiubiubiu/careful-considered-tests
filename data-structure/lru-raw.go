package main

import "fmt"

type ListNode struct {
	Key, Value int
	Pre, Next  *ListNode
}

type LinkedList struct {
	begin, tail *ListNode
}

func NewLinkedList() *LinkedList {
	list := new(LinkedList)
	list.begin = new(ListNode)
	list.tail = new(ListNode)
	list.tail.Value = -1

	list.begin.Next = list.tail
	list.tail.Pre = list.begin

	return list
}

func (l *LinkedList) deleteNode(node *ListNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (l *LinkedList) insertToHead(node *ListNode) {
	if node == nil {
		return
	}
	original := l.begin.Next

	l.begin.Next = node
	node.Pre = l.begin
	node.Next = original
	original.Pre = node
}

// AddNodeToHead 新建node，并插入到头结点
func (l *LinkedList) AddNodeToHead(key, value int) *ListNode {
	node := &ListNode{
		Key:   key,
		Value: value,
	}

	l.insertToHead(node)
	return node
}

// MoveNodeToHead 将节点从其他位置，插入到头结点
func (l *LinkedList) MoveNodeToHead(node *ListNode) {
	if node == nil {
		return
	}
	l.deleteNode(node)
	l.insertToHead(node)
}

type LRURaw struct {
	list *LinkedList
	m    map[int]*ListNode
	cap  int
}

func NewLRURaw(cap int) *LRURaw {
	return &LRURaw{
		list: NewLinkedList(),
		m:    make(map[int]*ListNode, cap),
		cap:  cap,
	}
}

func (lru *LRURaw) IsUnderCap() bool {
	return len(lru.m) < lru.cap
}

func (lru *LRURaw) Get(key int) (value int, ok bool) {
	node, isExist := lru.m[key]
	if !isExist {
		return 0, false
	}
	// move node to list head
	lru.list.MoveNodeToHead(node)
	return node.Value, true
}

func (lru *LRURaw) Set(key, value int) {
	node, isExist := lru.m[key]
	if !isExist {
		if !lru.IsUnderCap() {
			lru.list.deleteNode(lru.list.tail.Pre)
			delete(lru.m, lru.list.tail.Pre.Key)
		}
		node = lru.list.AddNodeToHead(key, value)
		lru.m[key] = node
		return
	}
	// move node to list head
	node.Value = value
	lru.list.MoveNodeToHead(node)
	return
}

func main() {
	lru := NewLRURaw(5)
	for i := 1; i <= 10; i++ {
		lru.Set(i, i)
	}
	fmt.Println("map size", len(lru.m))

	fmt.Println("ok")
}
