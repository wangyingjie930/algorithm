/**
  @author: wangyingjie
  @since: 2023/4/3
  @desc: https://leetcode.cn/problems/palindrome-linked-list/description/
**/

package 回文链表

import (
	List "algorithm/链表/单链表"
)

func isPalindrome(head *List.Node) bool {
	//获取链表中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	//反转后半部分链表
	left, right := head, reverse(slow)

	//两边同时进行遍历
	for left != nil && right != nil && left.Val == right.Val {
		left = left.Next
		right = right.Next
	}

	//当右边的链表遍历完即完成
	return right == nil
}

func reverse(head *List.Node) *List.Node {
	var p *List.Node
	for head != nil {
		tmp := head.Next
		head.Next = p
		p = head
		head = tmp
	}
	return p
}
