/**
  @author: wangyingjie
  @since: 2023/3/22
  @desc: https://leetcode.cn/problems/sort-list/description/
  掌握程度: 不熟
**/

package 基于归并排序

import (
	List "algorithm/链表/单链表"
)

func sortList(head *List.Node) *List.Node {
	return helper(head, nil)
}

// helper
//  @Description: 返回的是[head,tail)排完序之后的链表
//  @param head
//  @param tail
//  @return *List.Node
func helper(head, tail *List.Node) *List.Node {
	if head == tail || head.Next == tail {
		return head
	}

	var mid *List.Node
	head, mid = partition(head, tail) //分区之后, 原有的链表头结点变动, 需要重新找头结点
	ret := helper(head, mid)          // 返回的是待排序分区的[head, mid)链表
	mid.Next = helper(mid.Next, tail) // 返回的是待排序分区的[mid.Next, tail)链表, 并且需要将两个链表合并
	return ret
}

// partition
//  @Description: 比较可以理解的写法, 维护两个链表, 一边存放小于head, 一边存放大于head, 之后进行合并
//  @param head
//  @param tail
//  @return *List.Node
//  @return *List.Node
func partition(head, tail *List.Node) (*List.Node, *List.Node) {
	left := new(List.Node)
	right := new(List.Node)
	l, r, cur := left, right, head.Next
	for cur != tail {
		tmp := cur.Next
		cur.Next = nil //处理之前先断开之前的连接

		if cur.Val < head.Val {
			//小于head放到left链表
			l.Next = cur
			l = l.Next
		} else {
			//大于head放到right链表
			r.Next = cur
			r = r.Next
		}
		cur = tmp
	}

	//重点
	r.Next = tail          // 右边的链表要重新连接原有链表的tail, 否则会失联
	head.Next = right.Next //head右边连接right链表
	l.Next = head          //左边连接left链表
	return left.Next, head //返回新的链表头结点, 分区之后的节点
}

/*
解法比较难理解
func partition(head, tail *List.Node) (*List.Node, *List.Node) {
	lhead, rhead, cur := head, head, head.Next
	for cur != tail {
		next := cur.Next
		if cur.Val < head.Val {
			cur.Next = lhead
			lhead = cur
		} else {
			rhead.Next = cur
			rhead = cur
		}
		cur = next
	}
	rhead.Next = tail
	return lhead, head
}*/
