/**
  @author: wangyingjie
  @since: 2023/3/26
  @desc: https://leetcode.cn/problems/balanced-binary-tree/
**/

package 平衡二叉树

import (
	"algorithm/二叉查找树/Tree"
	"math"
)

func isBalanced(root *Tree.Node) bool {
	if root == nil {
		return true
	}

	return math.Abs(float64(maxDepth(root.Right))-float64(maxDepth(root.Left))) < 1 &&
		isBalanced(root.Right) && isBalanced(root.Left)
}

func maxDepth(root *Tree.Node) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
