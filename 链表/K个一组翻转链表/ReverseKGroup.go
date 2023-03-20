/**
  @author: wangyingjie
  @since: 2023/2/8
  @desc: K 个一组翻转链表
  @see: https://leetcode.cn/problems/reverse-nodes-in-k-group/
  掌握程度: 边界处理掌握不够 递归+迭代
**/

package K个一组翻转链表

import List "algorithm/链表/单链表"

func reverseKGroup(head *List.Node, k int) *List.Node {
	tailNext := head

	for i := 0; i < k; i++ {
		if tailNext == nil {
			//剩下不满足的无需反转直接返回
			return head
		}
		tailNext = tailNext.Next
	}

	reverseList := reverse(head, tailNext)
	newList := reverseKGroup(tailNext, k)
	head.Next = newList
	return reverseList
}

// reverse
//  @Description: 普通链表反转
//  @param head
//  @param tailNext 注意是尾结点的下一个节点
//  @return *List.Node
func reverse(head, tailNext *List.Node) *List.Node {
	var pre *List.Node
	for head != tailNext {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	return pre
}

// reverseKGroupIteration
//  @Description: 使用迭代法实现
//  @param head
//  @param k
//  @return *List.Node
func reverseKGroupIteration(head *List.Node, k int) *List.Node {
	dummyNode := &List.Node{Next: head} //关键: 因为中间反转后, 需要有一个pre和next来进行串联
	pre := dummyNode                    //这里就是需要串联的开头
	for head != nil {
		tail := head
		for i := 1; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummyNode.Next
			}
		}
		tailNext := tail.Next

		reverseList := reverse(head, tailNext)
		pre.Next = reverseList //头串联
		head.Next = tailNext   //尾串联

		pre = head
		head = pre.Next
	}
	return dummyNode.Next
}
