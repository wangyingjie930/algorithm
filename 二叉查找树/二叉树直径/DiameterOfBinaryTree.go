/**
  @author: wangyingjie
  @since: 2023/2/16
  @desc: https://leetcode.cn/problems/diameter-of-binary-tree/description/
**/

package 二叉树直径

import (
	"algorithm/二叉查找树/Tree"
)

// diameterOfBinaryTree
//  @Description: 就是在求深度的时候顺便将直径求出来
// 深度: 表示从根节点到叶节点完整的节点个数
// 左节点深度 + 右节点深度 + 1 = 一共经历的节点数 = 直径+1
//  @param root
//  @return int
func diameterOfBinaryTree(root *Tree.Node) int {
	maxDistance := 0 //经历的节点数
	height(root, &maxDistance)
	return maxDistance - 1 //节点数-1等于直径
}

func height(root *Tree.Node, maxDistance *int) int {
	if root == nil {
		return 0
	}
	left := height(root.Left, maxDistance)
	right := height(root.Right, maxDistance)
	*maxDistance = max(*maxDistance, left+right+1)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
