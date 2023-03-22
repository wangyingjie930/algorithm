/**
  @author: wangyingjie
  @since: 2023/3/22
  @desc: https://leetcode.cn/problems/add-two-numbers/description/
**/

package 两数相加

import List "algorithm/链表/单链表"

func addTwoNumbers(l1 *List.Node, l2 *List.Node) *List.Node {
	add := 0   //保存进位信息
	head := l1 //将结果统一存放在l1链表
	var lastL1 *List.Node

	for l1 != nil && l2 != nil {
		//按位加法
		res := l1.Val + l2.Val + add
		l1.Val = res % 10
		add = res / 10

		lastL1 = l1
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 == nil && lastL1 != nil {
		//l2比较长, 转移到l1中
		lastL1.Next = l2
		l1 = l2 //指向待补充进位的位数
	}

	for l1 != nil {
		//对于进位信息的维护
		res := l1.Val + add
		l1.Val = res % 10
		add = res / 10
		lastL1 = l1
		l1 = l1.Next
	}

	if add > 0 && lastL1 != nil {
		lastL1.Next = &List.Node{Val: add, Next: nil}
	}

	return head
}
