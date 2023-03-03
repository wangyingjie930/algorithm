/**
  @author: wangyingjie
  @since: 2023/3/3
  @desc: https://leetcode.cn/problems/binary-tree-right-side-view/
**/

package 二叉树的右视图

import "algorithm/二叉查找树/Tree"

// rightSideView
//  @Description: 实际就是层序遍历, 取每一层的最右节点
//  @param root
//  @return []int
func rightSideView(root *Tree.Node) []int {
	if root == nil {
		return []int{}
	}
	var queue []*Tree.Node
	childQueue := []*Tree.Node{root}
	var ret []int

	for len(childQueue) > 0 {
		ret = append(ret, childQueue[len(childQueue)-1].Value)
		queue = childQueue
		childQueue = []*Tree.Node{}

		for len(queue) > 0 {
			popNode := queue[0]
			queue = queue[1:]
			if popNode.Left != nil {
				childQueue = append(childQueue, popNode.Left)
			}
			if popNode.Right != nil {
				childQueue = append(childQueue, popNode.Right)
			}
		}
	}
	return ret
}
