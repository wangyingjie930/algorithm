package 链式队列

type Node struct {
	Val string
	Next *Node
}

type LinkQueue struct {
	head *Node
	tail *Node
	len int
}

func NewLinkQueue() *LinkQueue {
	return &LinkQueue{len: 0}
}

func (linkQueue *LinkQueue) enQueue(val string) {
	node := &Node{Val: val}
	if linkQueue.len == 0 {
		linkQueue.head = node
		linkQueue.tail = node
	}else {
		linkQueue.tail.Next = node
		linkQueue.tail = linkQueue.tail.Next
	}
	linkQueue.len++
}

func (linkQueue *LinkQueue) deQueue() string {
	if linkQueue.len == 0 {
		return ""
	}else {
		ret := linkQueue.head.Val
		linkQueue.head = linkQueue.head.Next
		linkQueue.len--
		return ret
	}
}

