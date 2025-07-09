package main

import (
	"fmt"
	"math"
)

func main() {
	lru := InitLru(3)
	lru.Put("a", 1)
	lru.Put("b", 2)
	lru.Put("c", 3)
	fmt.Println(lru.Head())
	fmt.Println(lru.Iterator())
	lru.Put("d", 4)
	fmt.Println(lru.Iterator())
}

func InitLru(size int) *Lru {
	head := &LruNode{}
	tail := &LruNode{
		Value: math.MinInt32,
	}
	head.Next = tail
	tail.Prev = head
	return &Lru{
		Size:    size,
		nodeMap: make(map[string]*LruNode),
		head:    head,
		tail:    tail,
	}
}

type Lru struct {
	Size    int
	nodeMap map[string]*LruNode
	head    *LruNode
	tail    *LruNode
}

func (r *Lru) Iterator() []int {
	curNode := r.head.Next
	res := make([]int, 0, len(r.nodeMap))
	for curNode != nil {
		if curNode.Value == math.MinInt32 {
			break
		}
		res = append(res, curNode.Value)
		curNode = curNode.Next
	}
	return res
}

func (r *Lru) Put(key string, value int) {
	_, has := r.nodeMap[key]
	if has {
		node := r.nodeMap[key]
		node.Value = value
		r.Get(key)
		return
	}
	node := &LruNode{
		Key:   key,
		Value: value,
	}
	if len(r.nodeMap) == r.Size {
		lastNode := r.tail.Prev
		prevNode := lastNode.Prev
		prevNode.Next = r.tail
		r.tail.Prev = prevNode
		delete(r.nodeMap, lastNode.Key)
	}
	nextNode := r.head.Next
	r.head.Next = node
	node.Prev = r.head
	node.Next = nextNode
	if nextNode != nil {
		nextNode.Prev = node
	}
	r.nodeMap[key] = node
}

func (r *Lru) Get(key string) int {
	node, has := r.nodeMap[key]
	if !has {
		return -1
	}
	prevNode := node.Prev
	nextNode := node.Next
	prevNode.Next = nextNode
	nextNode.Prev = prevNode

	headNext := r.head.Next
	r.head.Next = node
	node.Prev = r.head
	node.Next = headNext
	headNext.Prev = node
	return node.Value
}

func (r *Lru) Head() int {
	if r.Size <= 0 {
		return -1
	}
	return r.head.Next.Value
}

type LruNode struct {
	Key   string
	Value int
	Prev  *LruNode
	Next  *LruNode
}
