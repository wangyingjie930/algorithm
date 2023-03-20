/**
  @author: wangyingjie
  @since: 2023/3/21
  @desc: https://leetcode.cn/problems/binary-tree-maximum-path-sum/description/
  掌握程度: 需要重写
**/

package 二叉树中的最大路径和

import (
	"algorithm/二叉查找树/Tree"
	"math"
)

func maxPathSum(root *Tree.Node) int {
	maxSum := math.MinInt32
	dfs(root, &maxSum)
	return maxSum
}

// dfs
//  @Description: 代表走到root时, 能增加多少收益
//  @param root
//  @return int
func dfs(root *Tree.Node, maxSum *int) int {
	if root == nil {
		return 0
	}
	left := max(dfs(root.Left, maxSum), 0)   //往左边走能增加多少, 如果是负数表示收益降低, 取0, 表示切断左边的路不走了
	right := max(dfs(root.Right, maxSum), 0) //往右边走能增加多少, 同理

	innerMaxSum := left + root.Value + right
	*maxSum = max(*maxSum, innerMaxSum) //左边收益+右边收益+加当前收益取最大值, 目的是求以root为根的完整路径是否为最大路径

	//root只能一条路走到底, 所以使用max表示选左右两边哪一边的路径最大.
	//目的是给root的父亲节点提供root身为父节点的部分路径能为它增加多少收益
	return root.Value + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
