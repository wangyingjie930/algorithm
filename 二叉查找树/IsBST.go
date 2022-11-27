package BST

import "math"

/**
使用递归实现的, 但是空间复杂度不太行..
 */
func IsValidBST(head *Node) bool {
	if head == nil {
		return true
	}
	leftIsBST := IsValidBST(head.Left)
	rightIsBST := IsValidBST(head.Right)
	if leftIsBST && rightIsBST &&
		(head.Left == nil || head.Key > head.Left.Key) &&
		(head.Right == nil || head.Key < head.Right.Key) {
		return true
	}
	return false
}

/**
使用中序遍历, 如果并没有从小到大进行输出则不是一棵二叉查找树
 */
func IsValidBST2(head *Node) bool {
	if head == nil {
		return true
	}
	var stack []*Node
	tmp := math.MinInt64
	for len(stack) > 0 || head != nil {
		if head != nil {
			stack = append(stack, head)
			head = head.Left
		}else {
			head = stack[len(stack) - 1]
			stack = stack[:(len(stack) - 1)]
			if tmp > head.Key {
				return false
			}else {
				tmp = head.Key
			}
			head = head.Right
		}
	}
	return true
}
