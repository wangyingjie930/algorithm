/**
  @author: wangyingjie
  @since: 2023/3/20
  @desc: https://leetcode.cn/problems/binary-tree-preorder-traversal/
**/

package 二叉树的前序遍历

import "algorithm/二叉查找树/Tree"

type preorder interface {
	preorderTraversal(root *Tree.Node) []int
}

// Recur 使用递归
type Recur struct{}

func (Recur) preorderTraversal(root *Tree.Node) []int {
	var ret []int
	helper(root, &ret)
	return ret
}
func helper(root *Tree.Node, ret *[]int) {
	if root == nil {
		return
	}
	*ret = append(*ret, root.Value)
	helper(root.Left, ret)
	helper(root.Right, ret)
}

// Iteration 使用迭代
type Iteration struct{}

func (Iteration) preorderTraversal(root *Tree.Node) []int {
	var stack []*Tree.Node
	var ret []int
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			ret = append(ret, root.Value)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		root = node.Right
	}
	return ret
}

// Morris 遍历
type Morris struct{}

func (m Morris) preorderTraversal(root *Tree.Node) []int {
	var ret []int
	for root != nil {
		if root.Left != nil {
			right := root.Left
			//找到左子树中的最右节点
			for right.Right != nil && right.Right != root {
				right = right.Right
			}
			if right.Right == nil {
				ret = append(ret, root.Value)
				//表示第一次进来, 创建线索
				right.Right = root
				//向左子树方向
				root = root.Left
			} else {
				//线索还原
				right.Right = nil
				//向右子树前进
				root = root.Right
			}
		} else {
			ret = append(ret, root.Value)
			//包含操作: 到底部尝试向上走; 或者是左子树没有节点, 向右子树前进
			root = root.Right
		}
	}
	return ret
}
