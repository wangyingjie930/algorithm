package List

/**
两条单链表是否相交
1. 分别遍历两条链表, 分别获得两条结尾指针, 若相等, 则证明相交
2. 获得两者长度之差, 长的先走这几不差值, 然后一起走n步, 第一次相同的节点为相交节点
 */
func GetIntersectionNode(headA, headB *Node) *Node {
	if headA == nil {
		return nil
	}
	if headB == nil {
		 return nil
	}
	alen := 1
	aEnd := headA
	for aEnd.Next != nil {
		alen ++
		aEnd = aEnd.Next
	}
	blen := 1
	bEnd := headB
	for bEnd.Next != nil {
		blen ++
		bEnd = bEnd.Next
	}

	if bEnd != aEnd {
		return nil
	}
	if blen > alen {
		 headA, headB = headB, headA
		 alen, blen = blen, alen
	}
	diff := alen - blen
	for headA != headB {
		if diff > 0 {
			 headA = headA.Next
			 diff --
			 continue
		}
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}
