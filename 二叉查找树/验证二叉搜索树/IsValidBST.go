/**
  @author: wangyingjie
  @since: 2023/2/17
  @desc: https://leetcode.cn/problems/validate-binary-search-tree/description/
**/

package 验证二叉搜索树

import (
	BST "algorithm/二叉查找树/Tree"
	"math"
)

func isValidBST(root *BST.Node) bool {
	return judgeNode(root, math.MinInt32, math.MaxInt32)
}

// judgeNode
//  @Description: 根据当前节点的值划分区间
//  @param root
//  @param lower
//  @param upper
//  @return bool
func judgeNode(root *BST.Node, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Value <= lower || root.Value >= upper {
		return false
	}

	left := judgeNode(root.Left, lower, root.Value)
	right := judgeNode(root.Right, root.Value, upper)
	return left && right
}

/**
使用中序遍历, 如果并没有从小到大进行输出则不是一棵二叉查找树
*/
func IsValidBST2(head *BST.Node) bool {
	if head == nil {
		return true
	}
	var stack []*BST.Node
	tmp := math.MinInt64
	for len(stack) > 0 || head != nil {
		if head != nil {
			stack = append(stack, head)
			head = head.Left
		} else {
			head = stack[len(stack)-1]
			stack = stack[:(len(stack) - 1)]
			if tmp > head.Key {
				return false
			} else {
				tmp = head.Key
			}
			head = head.Right
		}
	}
	return true
}
