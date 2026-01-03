package main

type DoubleListNode struct {
	Key  int
	Val  int
	Next *DoubleListNode
	Prev *DoubleListNode
}

type LRUCache struct {
	Cap  int
	data map[int]*DoubleListNode
	MRU  *DoubleListNode
	LRU  *DoubleListNode
}

func ConstructorLRU(capacity int) LRUCache {
	obj := LRUCache{Cap: capacity,
		data: make(map[int]*DoubleListNode),
		MRU:  &DoubleListNode{Key: -1, Val: -1},
		LRU:  &DoubleListNode{Key: -1, Val: -1},
	}
	obj.LRU.Next, obj.MRU.Prev = obj.MRU, obj.LRU
	return obj
}

func (this *LRUCache) remove(node *DoubleListNode) {
	prev, next := node.Prev, node.Next
	prev.Next, next.Prev = next, prev
}

func (this *LRUCache) insert(node *DoubleListNode) {
	prev, nxt := this.MRU.Prev, this.MRU
	prev.Next = node
	nxt.Prev = node
	node.Prev, node.Next = prev, nxt
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.data[key]; ok {
		this.remove(node)
		this.insert(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.data[key]; ok {
		this.remove(node)
	}
	node := &DoubleListNode{Key: key, Val: value}
	this.data[key] = node
	this.insert(node)
	if len(this.data) > this.Cap {
		lru := this.LRU.Next
		this.remove(lru)
		delete(this.data, lru.Key)
	}
}
