package 单链表

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func NewSingleList() *List {
	return &List{Head: nil, Tail: nil}
}

// AddNode
//  @Description: 添加节点
//  @receiver list
//  @param val
func (list *List) AddNode(val int) {
	node := &Node{Val: val}
	if list.Head == nil {
		list.Head = node
		list.Tail = node
	} else {
		list.Tail.Next = node
		list.Tail = node
	}
}

// Traverse
//  @Description: 遍历链表
//  @receiver list
func (list *List) Traverse() {
	if list.Head == nil {
		fmt.Println("-> 空链表!")
		return
	}
	cur := list.Head
	for cur != nil {
		fmt.Printf("%d -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

//删除第k个结点的函数
func (list *List) deleteKth(k int) {
	if list.Head == nil {
		return
	}

	//设置哨兵, 带头节点
	sentinel := &Node{Val: -1, Next: list.Head}
	list.Head = sentinel
	pre := list.Head

	//找到第k-1个节点
	for i := 1; i <= k-1; i++ {
		pre = pre.Next
	}
	//进行删除
	pre.Next = pre.Next.Next
	//还原回原来的不带头节点
	list.Head = sentinel.Next
}

/**
设置两个指针, 两者间隔k个节点, 当快的指针遍历到nil之后
后面慢的指针正好处于倒数k个节点的前一个
*/
func (list *List) removeNthFromEnd(k int) {
	//这里是带了个哨兵, 防止出现链表只有一个节点的情况,
	//影响理解后续的代码实现逻辑这里可以忽略不计
	sentinel := &Node{Val: -1, Next: list.Head}
	list.Head = sentinel

	cur := list.Head
	for ; cur != nil; cur = cur.Next {
		quick := cur
		i := 0
		//quick保持与cur间隔为k个节点
		for ; i <= k && quick != nil; i++ {
			quick = quick.Next
		}
		//quick此时遍历完成, cur处于倒数k个节点的前一个节点
		if quick == nil {
			//进行删除操作
			cur.Next = cur.Next.Next
			break
		}
	}
	list.Head = sentinel.Next
}
