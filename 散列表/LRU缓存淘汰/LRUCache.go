package LRU缓存淘汰

import "fmt"

// Node 双向链表节点
type Node struct {
	Prev, Next *Node
	Key int
	Value int
}

func (node *Node) String() string {
	return fmt.Sprintf("Key: %d, Value: %d", node.Key, node.Value)
}

// LRUCache 哈希表
type LRUCache struct {
	cache map[int]*Node
	count int
	cap int
	head, tail *Node
}

func (this LRUCache) String() string {
	h1 := this.head.Next
	ret := ""
	for h1 != this.tail {
		ret = fmt.Sprintf("%s->{%s}", ret, h1)
		h1 = h1.Next
	}
	return ret
}


func Constructor(capacity int) LRUCache {
	head := &Node{Key: 0, Value: 0}
	tail := &Node{Key: 0, Value: 0}
	head.Next = tail
	tail.Prev = head
	return LRUCache{make(map[int]*Node, capacity), 0, capacity, head, tail}
}


func (this *LRUCache) Get(key int) int {
	if this.cache[key] == nil {
		return -1
	}
	node := this.cache[key]
	//将节点移至双向链表之前
	this.moveToHead(node)
	return node.Value
}


func (this *LRUCache) Put(key int, value int)  {
	if this.cache[key] == nil {
		//如果不存在
		node := &Node{Key: key, Value: value}
			//如果容量不够, 删除最后的节点
		if this.cap == this.count {
			this.deleteNode(this.tail.Prev)
		}
			//如果容量够, 将node添加到头部
		this.addToHead(node)
		this.cache[key] = node
	}else {
		//如果存在
		node := this.cache[key]
			//更新value
		node.Value = value
			//移动到头部
		this.moveToHead(node)
	}
}

// 将节点删除
func (this *LRUCache) deleteNode(deleteNode *Node)  {
	deleteNode.Prev.Next = deleteNode.Next
	deleteNode.Next.Prev = deleteNode.Prev

	delete(this.cache, deleteNode.Key)
	deleteNode = nil
	this.count --
}

// 移到到头部
func (this *LRUCache) moveToHead(node *Node)  {
	this.deleteNode(node)
	this.addToHead(node)
}

// 添加到头部
func (this *LRUCache) addToHead(node *Node)  {
	node.Next = this.head.Next
	this.head.Next.Prev = node
	this.head.Next = node
	node.Prev = this.head

	this.cache[node.Key] = node
	this.count ++
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
