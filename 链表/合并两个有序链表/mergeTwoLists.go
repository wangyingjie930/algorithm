/**
  @author: wangyingjie
  @since: 2023/3/18
  @desc: https://leetcode.cn/problems/merge-two-sorted-lists/description/
**/

package 合并两个有序链表

import List "algorithm/链表/单链表"

// mergeTwoLists
//  @Description: 指针遍历写法
//  @param list1
//  @param list2
//  @return *List.Node
func mergeTwoLists(list1 *List.Node, list2 *List.Node) *List.Node {
	list := &List.Node{Val: -1}
	header := list

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			list.Next = list1
			list1 = list1.Next
		} else {
			list.Next = list2
			list2 = list2.Next
		}
		list = list.Next
	}

	if list1 != nil {
		list.Next = list1
	}
	if list2 != nil {
		list.Next = list2
	}

	return header.Next
}

// mergeTwoListsRecursion
//  @Description: 递归写法
//  @param list1
//  @param list2
//  @return *List.Node
func mergeTwoListsRecursion(list1 *List.Node, list2 *List.Node) *List.Node {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoListsRecursion(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoListsRecursion(list1, list2.Next)
		return list2
	}
}

func mergeTwoListsDistinct(list1 *List.Node, list2 *List.Node) *List.Node {
	list := &List.Node{Val: -1}
	header := list

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			list.Next = list1
			list1 = list1.Next
		} else {
			list.Next = list2
			list2 = list2.Next
		}
		list = list.Next
	}

	if list1 != nil {
		list.Next = list1
	}
	if list2 != nil {
		list.Next = list2
	}

	return header.Next
}
