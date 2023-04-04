/**
  @author: wangyingjie
  @since: 2023/4/3
  @desc: https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
**/

package 链表中倒数第k个节点

import List "algorithm/链表/单链表"

func getKthFromEnd(head *List.Node, k int) *List.Node {
	slow, fast := head, head
	for k > 0 && fast != nil {
		fast = fast.Next
		k--
	}
	if k != 0 {
		return nil
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
