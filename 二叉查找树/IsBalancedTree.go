package BST

import (
	"algorithm/二叉查找树/Tree"
	"math"
)

/**
判断是否为平衡二叉树
1. 判断左树是否为平衡, 高度为多少
2. 判断右树是否为平衡, 高度为多少
3. abs(左高-右高) <= 1 则为平衡
*/
func IsBalance(head *Tree.Node) (bool, int) {
	if head == nil {
		return true, 0
	}

	leftBalance, leftHeight := IsBalance(head.Left)
	rightBalance, rightHeight := IsBalance(head.Right)

	if leftBalance && rightBalance && math.Abs(float64(leftHeight)-float64(rightHeight)) <= 1 {
		return true, int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
	}
	return false, int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}
