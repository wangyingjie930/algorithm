/**
  @author: wangyingjie
  @since: 2023/3/18
  @desc: https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
**/

package 删除排序链表中的重复元素

import List "algorithm/链表/单链表"

func deleteDuplicates(head *List.Node) *List.Node {
	slow := head
	for slow != nil {
		quick := slow.Next
		for quick != nil && quick.Val == slow.Val {
			quick = quick.Next
		}
		slow.Next = quick
		slow = slow.Next
	}
	return head
}
