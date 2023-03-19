/**
  @author: wangyingjie
  @since: 2023/3/20
  @desc: https://leetcode.cn/problems/binary-tree-level-order-traversal/description/
**/

package 二叉树的层序遍历

import "algorithm/二叉查找树/Tree"

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
