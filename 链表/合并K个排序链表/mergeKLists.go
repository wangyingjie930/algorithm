/**
  @author: wangyingjie
  @since: 2023/3/10
  @desc: https://leetcode.cn/problems/merge-k-sorted-lists/solutions/
  使用归并+链表合并
**/

package 合并K个排序链表

import (
	List "algorithm/链表/单链表"
)

func mergeKLists(lists []*List.Node) *List.Node {
	if len(lists) == 0 {
		return nil
	}
	return helper(lists, 0, len(lists)-1)
}

// helper
//  @Description: 归并
//  @param lists
//  @param start
//  @param end
//  @return *List.Node
func helper(lists []*List.Node, start, end int) *List.Node {
	if start >= end {
		return lists[end]
	}

	mid := (start + end) / 2
	l1 := helper(lists, start, mid)
	l2 := helper(lists, mid+1, end)
	return mergeTwoList(l1, l2)
}

// mergeTwoList
//  @Description: 合并两条链表
//  @param l1
//  @param l2
//  @return *List.Node
func mergeTwoList(l1 *List.Node, l2 *List.Node) *List.Node {
	dummy := &List.Node{}
	p := dummy
	for l1 != nil && l2 != nil {
		if l2.Val < l1.Val {
			p.Next = l2
			l2 = l2.Next
		} else {
			p.Next = l1
			l1 = l1.Next
		}
		p = p.Next
	}

	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}

	return dummy.Next
}
