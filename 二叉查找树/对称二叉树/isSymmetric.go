/**
  @author: wangyingjie
  @since: 2023/4/2
  @desc: https://leetcode.cn/problems/symmetric-tree/description/?orderBy=most_votes
**/

package 对称二叉树

import "algorithm/二叉查找树/Tree"

func isSymmetric(root *Tree.Node) bool {
	return check(root, root)
}

// check
//  @Description: 重点是能想出同时分开进行对比
//  @param a
//  @param b
//  @return bool
func check(a, b *Tree.Node) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	}

	return a.Value == b.Value && check(a.Left, b.Right) && check(a.Right, b.Left)
}
