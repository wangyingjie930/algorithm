package 单链表

import "fmt"

type Node struct {
	Val int
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

func NewSingleList() *List {
	return &List{head: nil, tail: nil}
}

//添加节点
func (list *List) addNode(val int) {
	node := &Node{Val: val}
	if list.head == nil {
		list.head = node
		list.tail = node
	}else {
		list.tail.next = node
		list.tail = node
	}
}

// 遍历链表
func (list *List) traverse() {
	if list.head == nil {
		fmt.Println("-> 空链表!")
		return
	}
	cur := list.head
	for cur != nil {
		fmt.Printf("%d -> ", cur.Val)
		cur = cur.next
	}
	fmt.Println()
}

//删除第k个结点的函数
func (list *List) deleteKth(k int) {
	if list.head == nil {
		return
	}

	//设置哨兵, 带头节点
	sentinel := &Node{Val: -1, next: list.head}
	list.head = sentinel
	pre := list.head

	//找到第k-1个节点
	for i := 1; i <= k - 1; i ++ {
		pre = pre.next
	}
	//进行删除
	pre.next = pre.next.next
	//还原回原来的不带头节点
	list.head = sentinel.next
}


/**
设置两个指针, 两者间隔k个节点, 当快的指针遍历到nil之后
后面慢的指针正好处于倒数k个节点的前一个
 */
func (list *List) removeNthFromEnd(k int) {
	//这里是带了个哨兵, 防止出现链表只有一个节点的情况,
	//影响理解后续的代码实现逻辑这里可以忽略不计
	sentinel := &Node{Val: -1, next: list.head}
	list.head = sentinel

	cur := list.head
	for ;cur != nil; cur = cur.next {
		quick := cur
		i := 0
		//quick保持与cur间隔为k个节点
		for ; i <= k  && quick != nil; i++ {
			quick = quick.next
		}
		//quick此时遍历完成, cur处于倒数k个节点的前一个节点
		if quick == nil{
			//进行删除操作
			cur.next = cur.next.next
			break
		}
	}
	list.head = sentinel.next
}
