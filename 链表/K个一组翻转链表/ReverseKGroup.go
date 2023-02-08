/**
  @author: wangyingjie
  @since: 2023/2/8
  @desc: K 个一组翻转链表
  @see: https://leetcode.cn/problems/reverse-nodes-in-k-group/
**/

package K个一组翻转链表

import (
	SingList "algorithm/链表/单链表"
	"errors"
)

func reverseKGroup(head *SingList.Node, k int) *SingList.Node {
	// 对第一组进行翻转以获得新的头指针
	h, cur, err := reverseList(head, k)
	if err != nil {
		return head
	}

	head = h
	for cur.Next != nil {
		//对下一组进行翻转, 返回头指针, 尾指针
		h, t, err := reverseList(cur.Next, k)
		if err != nil {
			break
		}
		//将前面一组的最后一个节点指向下一个组翻转后的的头节点
		cur.Next = h
		cur = t
	}
	return head
}

// reverseList
//  @Description: 进行链表的翻转操作
//  @param start
//  @param count
//  @return *SingList.Node
func reverseList(start *SingList.Node, count int) (head, tail *SingList.Node, err error) {
	if !canReverse(start, count) {
		return nil, nil, errors.New("不能进行链表的翻转操作")
	}

	cur := start
	var p *SingList.Node
	for i := count; i >= 1; i-- {
		next := cur.Next
		cur.Next = p
		p = cur
		cur = next
	}
	// 反转后, 之前的头结点为反转后的尾结点, 需要将start节点的下一个节点指向下一组还没反转的头节点
	start.Next = cur
	return p, start, nil
}

// canReverse
//  @Description: 判断是否能进行翻转
//  @param start
//  @param count
//  @return bool
func canReverse(start *SingList.Node, count int) bool {
	i := count
	for p := start; i >= 1 && p != nil; i-- {
		p = p.Next
	}
	return i == 0
}
