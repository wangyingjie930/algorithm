/**
  @author: wangyingjie
  @since: 2023/3/22
  @desc: https://leetcode.cn/problems/sort-list/description/
  掌握程度: 不熟
**/

package 基于归并排序

import List "algorithm/链表/单链表"

func sortList(head *List.Node) *List.Node {
	return helper(head, nil)
}

func helper(start, end *List.Node) *List.Node {
	if start == nil {
		return start
	}

	//当start.Next == end表示已经走到尾结点了(相当于普通链表最后节点.Next=nil), 此时需要断开end
	// 因为合并链表的前提是两个链表是分开的, 不是处在同一链表中, 只有两条链表的.next=nil才能进行合并
	if start.Next == end {
		start.Next = nil
		return start
	}

	//找到中点
	slow := start
	for quick := slow; quick != end; {
		quick = quick.Next
		if quick != end {
			slow = slow.Next
			quick = quick.Next
		}
	}

	return mergeList(helper(start, slow), helper(slow, end))
}

// mergeList
//  @Description: 合并两个有序链表
//  @param list1
//  @param list2
//  @return *List.Node
func mergeList(list1, list2 *List.Node) *List.Node {
	dummy := new(List.Node)
	p := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}
	return dummy.Next
}
