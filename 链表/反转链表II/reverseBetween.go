/**
  @author: wangyingjie
  @since: 2023/3/18
  @desc: https://leetcode.cn/problems/reverse-linked-list-ii/
  掌握程度: 需要题解辅助dummyNode
**/

package 反转链表II

import List "algorithm/链表/单链表"

func reverseBetween(head *List.Node, left int, right int) *List.Node {
	var preStart, afterEnd *List.Node            //关键: 要找出反转链表的前驱和后继
	dummyNode := &List.Node{Val: -1, Next: head} //关键:需要借助虚拟节点方便边界判断
	p := dummyNode
	for p != nil {
		left--
		if left == 0 {
			preStart = p
		}

		right--
		if right+2 == 0 {
			afterEnd = p
		}
		p = p.Next
	}
	start, end := reverse(preStart.Next, afterEnd)
	preStart.Next = start
	end.Next = afterEnd
	return dummyNode.Next
}

// reverse
//  @Description: 反转链表
//  @param start
//  @param end
//  @return *List.Node
//  @return *List.Node
func reverse(start, end *List.Node) (*List.Node, *List.Node) {
	p := start
	var pre *List.Node
	for p != end {
		tmp := p.Next
		p.Next = pre
		pre = p
		p = tmp
	}
	return pre, start
}

//todo: 使用头插法方式实现
