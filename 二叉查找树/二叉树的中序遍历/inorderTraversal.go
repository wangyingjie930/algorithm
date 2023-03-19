/**
  @author: wangyingjie
  @since: 2023/3/20
  @desc: https://leetcode.cn/problems/binary-tree-inorder-traversal/
**/

package 二叉树的中序遍历

import "algorithm/二叉查找树/Tree"

type inorder interface {
	inorderTraversal(root *Tree.Node) []int
}

type Recur struct{}

// inorderTraversal
//  @Description: 递归写法
//  @receiver Recur
//  @param root
//  @return []int
func (Recur) inorderTraversal(root *Tree.Node) []int {
	var ret []int
	helper(root, &ret)
	return ret
}
func helper(root *Tree.Node, ret *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, ret)
	*ret = append(*ret, root.Value)
	helper(root.Right, ret)
}

type Iteration struct{}

// inorderTraversal
//  @Description: 迭代写法
//  @param root
//  @return []int
func (Iteration) inorderTraversal(root *Tree.Node) []int {
	var stack []*Tree.Node
	var ret []int
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		ret = append(ret, node.Value)
		root = node.Right
	}
	return ret
}

type morris struct{}

func (morris) inorderTraversal(root *Tree.Node) []int {
	var res []int
	for root != nil {
		if root.Left != nil {
			//找出左子树中最右的节点
			right := root.Left
			for right.Right != nil && right.Right != root {
				right = right.Right
			}
			if right.Right == nil {
				//表示第一次来, 创建线索
				right.Right = root
				//因为没有来过, 所以尝试向左子树方向走
				root = root.Left
			} else {
				//表示来过, 还原线索
				right.Right = nil
				res = append(res, root.Value)
				//因为左子树来过, 所以尝试向右子树方向走
				root = root.Right
			}
		} else {
			res = append(res, root.Value)
			//当遍历到叶子节点, 这个的作用是快速回到上面的节点
			root = root.Right
		}
	}
	return res
}
