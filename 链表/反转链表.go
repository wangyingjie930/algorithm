package List

// ReverseList 使用3指针进行反转链表
func ReverseList(head *Node) *Node {
	if head == nil {
		return nil
	}
	var pre, cur *Node
	cur = head
	pre = nil
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
