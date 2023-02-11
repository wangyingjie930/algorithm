/**
  @author: wangyingjie
  @since: 2023/2/10
  @desc: https://leetcode.cn/problems/reorder-list/
**/

package 重排链表

import List "algorithm/链表/单链表"

func reorderList(head *List.Node) {
	// 寻找链表中点
	midNode := findMidNode(head)
	if midNode == head {
		return
	}
	// 反转后半段的链表
	midNode = reverseList(midNode)
	// 合并链表
	mergeList(head, midNode)
}

// findMidNode
//  @Description: 寻找链表的中点
//  @param head
//  @return *List.Node
func findMidNode(head *List.Node) *List.Node {
	cur := head
	quick := head
	var lastCur *List.Node
	for quick != nil && quick.Next != nil {
		lastCur = cur
		cur = cur.Next
		quick = quick.Next.Next
	}
	if lastCur != nil {
		lastCur.Next = nil
	}
	return cur
}

// reverseList
//  @Description: 反转链表
//  @param head
//  @return *List.Node
func reverseList(head *List.Node) *List.Node {
	var p *List.Node
	node := head
	for node != nil {
		next := node.Next
		node.Next = p
		p = node
		node = next
	}
	return p
}

// mergeList
//  @Description: 合并两个链表
//  @param l1
//  @param l2
func mergeList(l1 *List.Node, l2 *List.Node) {
	node1 := l1
	node2 := l2
	for node1.Next != nil && node2 != nil {
		nextNode1 := node1.Next
		nextNode2 := node2.Next
		node1.Next = node2
		node2.Next = nextNode1

		node1 = nextNode1
		node2 = nextNode2
	}
	if node2 != nil {
		node1.Next = node2
	}
}
