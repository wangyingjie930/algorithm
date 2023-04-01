/**
  @author: wangyingjie
  @since: 2023/4/2
  @desc: https://leetcode.cn/problems/swap-nodes-in-pairs/
**/

package 两两交换链表中的节点

import List "algorithm/链表/单链表"

// swapPairs
//  @Description: 递归
//  @param head
//  @return *List.Node
func swapPairs(head *List.Node) *List.Node {
	if head == nil || head.Next == nil {
		return head
	}

	newList := swapPairs(head.Next.Next)

	newHead := head.Next
	newHead.Next = head
	head.Next = newList
	return newHead
}

func swapPairsIteration(head *List.Node) *List.Node {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &List.Node{Next: head}
	pre := dummy
	for pre != nil && pre.Next != nil && pre.Next.Next != nil {
		tmp := pre.Next
		pre.Next = tmp.Next
		tmp.Next = pre.Next.Next
		pre.Next.Next = tmp

		pre = tmp
	}

	return dummy.Next
}
