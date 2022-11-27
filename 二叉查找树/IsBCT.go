package BST

/**
是否为一颗完全二叉搜索树

1, 任意节点有右无左, 返回false
2. 如果遇到第一个只有左没有右的节点, 那么后续的节点都为叶子节点
 */
func IsBCT(head *Node) bool  {
	if head == nil {
		return true
	}
	queue := []*Node{head}
	tmp := false
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Right != nil && node.Left == nil {
			return false
		}else if node.Left != nil && node.Right == nil {
			tmp = true
		}
		if tmp && (node.Right != nil || node.Left != nil) {
			return false
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return true
}
