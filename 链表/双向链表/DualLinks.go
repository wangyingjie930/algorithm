package 双向链表

type Node struct {
	Val int
	Next *Node
	Pre *Node
}

type List struct {
	head *Node
	tail *Node
}

//添加节点
func (list *List) addNode(val int) {
	node := &Node{Val: val}
	if list.head == nil {
		list.head = node
		list.tail = node
	}else {
		list.tail.Next = node
		node.Pre = list.tail
		list.tail = node
	}
}
