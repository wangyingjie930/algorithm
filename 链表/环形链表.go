package List

/**
获得入环时的第一个节点

思路就是快慢指针, 快指针走两步, 慢指针走一步
快慢相遇表示有环
相遇之后快指针返回头部, 变为走1步
再次相遇就是入环节点
 */
func DetectCycle(head *Node) *Node {
	slow := head
	fast := head
	isCycle := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			isCycle = true
			break
		}
	}
	if !isCycle {
		return nil
	}
	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}