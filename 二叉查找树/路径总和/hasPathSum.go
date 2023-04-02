/**
  @author: wangyingjie
  @since: 2023/4/2
  @desc: https://leetcode.cn/problems/path-sum/
**/

package 路径总和

import "algorithm/二叉查找树/Tree"

func hasPathSum(root *Tree.Node, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Value
	}
	return hasPathSum(root.Left, targetSum-root.Value) || hasPathSum(root.Right, targetSum-root.Value)
}
