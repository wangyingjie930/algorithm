package List

import "fmt"

type Node struct {
	Val int
	Next *Node
}

func NewList(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	head := &Node{arr[0], nil}
	p := head
	for _, v := range arr[1:] {
		node := &Node{v, nil}
		p.Next = node
		p = p.Next
	}
	return head
}

func ShowList(head *Node) {
	node := head
	for node != nil {
		fmt.Print(node.Val, ", ")
		node = node.Next
	}
	fmt.Println()
}

/**
合并两个有序链表
 */
func mergeTwoLists(list1 *Node, list2 *Node) *Node {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	//哨兵节点
	ret := &Node{Val: -1}
	cur := ret

	//双指针遍历
	for ; list1 != nil && list2 != nil; {
		if list1.Val > list2.Val {
			cur.Next = list2
			list2 = list2.Next
		}else {
			cur.Next = list1
			list1 = list1.Next
		}
		cur = cur.Next
	}

	if list2 != nil {
		cur.Next = list2
	}

	if list1 != nil {
		cur.Next = list1
	}

	return ret.Next
}

/**
链表的中间节点, 使用快慢指针
快指针到达的时候正是慢指针处于中间节点的时候
 */
func middleNode(head *Node) *Node {
	quick := head
	//慢指针head走一步
	for ; head != nil; {
		if quick.Next != nil && quick.Next.Next != nil {
			//快指针quick走两步
			quick = quick.Next.Next
		}else if quick.Next != nil {
			quick = quick.Next
		}else {
			break
		}
		head = head.Next
	}
	return head
}

/**
检测是否为环形链表
还是使用快慢指针, 如果相遇则证明是环形链表
 */
func hasCycle(head *Node) bool {
	quick := head
	isCycle := false
	//慢指针head走一步
	for ; head != nil; {
		if quick.Next != nil && quick.Next.Next != nil {
			//快指针quick走两步
			quick = quick.Next.Next
		}else {
			return false
		}
		if quick == head {
			isCycle = true
			break
		}
		head = head.Next
	}
	return isCycle
}

/**
合并k个有序链表
 */
func mergeKLists(lists []*Node) *Node {
	if len(lists) == 0 {
		return nil
	}
	ret := lists[0]
	for i := 1; i < len(lists); i++ {
		ret = mergeTwoLists(ret, lists[i])
	}
	return ret
}
