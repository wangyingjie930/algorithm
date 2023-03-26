/**
  @author: wangyingjie
  @since: 2023/3/26
  @desc: https://leetcode.cn/problems/maximum-depth-of-binary-tree/
**/

package 二叉树的最大深度

import "algorithm/二叉查找树/Tree"

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
