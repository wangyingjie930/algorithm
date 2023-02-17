/**
  @author: wangyingjie
  @since: 2023/2/17
  @desc: https://leetcode.cn/problems/intersection-of-two-linked-lists/
**/

package 相交链表

import List "algorithm/链表/单链表"

/**
两条单链表是否相交
1. 分别遍历两条链表, 分别获得两条结尾指针, 若相等, 则证明相交
2. 获得两者长度之差, 长的先走这几不差值, 然后一起走n步, 第一次相同的节点为相交节点
*/
func getIntersectionNode(headA, headB *List.Node) *List.Node {
	nodeA, nodeB := headA, headB
	lenA, lenB := 0, 0
	for nodeA.Next != nil {
		lenA++
		nodeA = nodeA.Next
	}
	for nodeB.Next != nil {
		lenB++
		nodeB = nodeB.Next
	}
	if nodeB != nodeA {
		//尾结点不在一起表示不相交
		return nil
	}

	nodeA, nodeB = headA, headB
	if lenB > lenA {
		//以nodeA作为最长的那条链表
		nodeA, nodeB = headB, headA
		lenA, lenB = lenB, lenA
	}

	for i := 0; i < lenA-lenB; i++ {
		//长的先走
		nodeA = nodeA.Next
	}
	for nodeA != nodeB {
		//一起走
		nodeA = nodeA.Next
		nodeB = nodeB.Next
	}
	return nodeA
}
