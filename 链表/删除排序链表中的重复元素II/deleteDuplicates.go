/**
  @author: wangyingjie
  @since: 2023/3/22
  @desc: https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/
  掌握程度: 不熟
**/

package 删除排序链表中的重复元素II

import List "algorithm/链表/单链表"

func deleteDuplicates(head *List.Node) *List.Node {
	dummy := &List.Node{Next: head}
	slow := dummy
	for slow != nil {
		if slow.Next != nil && slow.Next.Next != nil && slow.Next.Val == slow.Next.Next.Val {
			//发现重复
			quick := slow.Next
			for quick.Next != nil && quick.Val == quick.Next.Val {
				quick = quick.Next
			}
			slow.Next = quick.Next //slow只需删除对应重复的, 不需要移动, 因为可能移动后所处的位置不再进行判断
		} else {
			//没有重复
			slow = slow.Next
		}
	}
	return dummy.Next
}
