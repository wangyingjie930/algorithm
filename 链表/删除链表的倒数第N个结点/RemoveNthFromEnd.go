/**
  @author: wangyingjie
  @since: 2023/2/15
  @desc: https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/
**/

package 删除链表的倒数第N个结点

import List "algorithm/链表/单链表"

func removeNthFromEnd(head *List.Node, n int) *List.Node {
	dummy := &List.Node{Next: head}
	cur := dummy
	quick := dummy

	i := 0
	for ; i < n && quick != nil; i++ {
		quick = quick.Next
	}
	if i < n {
		return head
	}

	for quick != nil && quick.Next != nil {
		cur = cur.Next
		quick = quick.Next
	}
	//主要是这里
	cur.Next = cur.Next.Next
	return dummy.Next
}
