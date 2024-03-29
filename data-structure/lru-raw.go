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

	list.begin.Next = list.tail
	list.tail.Pre = list.begin

	return list
}

func (l *LinkedList) deleteNode(node *ListNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (l *LinkedList) AddNodeToHead(key, value int) {
	node := &ListNode{
		Key:   key,
		Value: value,
	}

	l.MoveNodeToHead(node)
}

func (l *LinkedList) MoveNodeToHead(node *ListNode) {
	if node == nil {
		return
	}
	originalHead := l.begin.Next
	node.Pre = l.begin
	l.begin.Next = node
	node.Next = originalHead
	originalHead.Pre = node
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
			delete(lru.m, key)
		}
		node := &ListNode{Key: key, Value: value}
		lru.list.MoveNodeToHead(node)
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
	for i := 0; i < 10; i++ {
		lru.Set(i, i)
	}

	fmt.Println("ok")
}
