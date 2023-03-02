/**
  @author: wangyingjie
  @since: 2023/3/2
  @desc: https://leetcode.cn/problems/reverse-linked-list/
**/

package 反转链表

import List "algorithm/链表/单链表"

// reverseList
//  @Description: 双指针
//  @param head
//  @return *List.Node
func reverseList(head *List.Node) *List.Node {
	i := head
	if head == nil {
		return nil
	}
	var p *List.Node
	for i != nil {
		tmp := i.Next
		i.Next = p
		p = i
		i = tmp
	}
	return p
}

// reverseListRecursion
//  @Description: 使用递归实现
//  @param head
//  @return *List.Node
func reverseListRecursion(head *List.Node) *List.Node {
	//1->2->3->4
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	// 1->2<-3<-4
	node := reverseListRecursion(head.Next)
	//1<-2<-3<-4
	head.Next.Next = head
	//nil<-1<-2<-3<-4
	head.Next = nil
	return node
}
