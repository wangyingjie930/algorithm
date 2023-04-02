/**
  @author: wangyingjie
  @since: 2023/4/2
  @desc: https://leetcode.cn/problems/invert-binary-tree/description/
**/

package 翻转二叉树

import "algorithm/二叉查找树/Tree"

func invertTree(root *Tree.Node) *Tree.Node {
	if root == nil {
		return nil
	}

	left := invertTree(root.Left)
	root.Left = invertTree(root.Right)
	root.Right = left

	return root
}
