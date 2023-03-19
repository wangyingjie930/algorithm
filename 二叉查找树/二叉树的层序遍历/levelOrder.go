/**
  @author: wangyingjie
  @since: 2023/3/20
  @desc: https://leetcode.cn/problems/binary-tree-level-order-traversal/description/
**/

package 二叉树的层序遍历

import "algorithm/二叉查找树/Tree"

// levelOrder
//  @Description: 使用队列实现
//  @param root
//  @return [][]int
func levelOrder(root *Tree.Node) [][]int {
	var queue []*Tree.Node
	tmpQueue := []*Tree.Node{root}
	var ret [][]int

	for len(tmpQueue) > 0 {
		queue = tmpQueue
		tmpQueue = []*Tree.Node{}

		var tmp []int
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				tmpQueue = append(tmpQueue, node.Left)
			}
			if node.Right != nil {
				tmpQueue = append(tmpQueue, node.Right)
			}
			tmp = append(tmp, node.Value)
		}
		ret = append(ret, tmp)
	}

	return ret
}

// levelOrderRecur
//  @Description: 使用递归实现层序遍历
//  @param root
//  @return [][]int
func levelOrderRecur(root *Tree.Node) [][]int {
	var ret [][]int
	helper(root, 0, &ret)
	return ret
}

// helper
//  @Description: 关键就是定义了level这个层数
//  @param root
//  @param level
//  @param ret
func helper(root *Tree.Node, level int, ret *[][]int) {
	if root == nil {
		return
	}
	if len(*ret)-1 < level {
		*ret = append(*ret, []int{})
	}
	(*ret)[level] = append((*ret)[level], root.Value)
	helper(root.Left, level+1, ret)
	helper(root.Right, level+1, ret)
}
