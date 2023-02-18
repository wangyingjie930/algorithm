/**
  @author: wangyingjie
  @since: 2023/2/18
  @desc: https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/
**/

package 二叉树的锯齿形层序遍历

import "algorithm/二叉查找树/Tree"

// zigzagLevelOrder
//  @Description: 核心点是采用一个双队列, 一个用来遍历当前的访问的节点, 另一个用来保存下一层的节点
//  @param root
//  @return [][]int
func zigzagLevelOrder(root *Tree.Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	var queue []*Tree.Node
	var childQueue []*Tree.Node
	childQueue = append(childQueue, root)
	var ret [][]int

	for i := 0; len(childQueue) > 0; i++ {
		//重置队列, 只遍历下一层的队列
		queue = childQueue
		//将下一层的队列清空, 用来容纳遍历过程中新的一层
		childQueue = []*Tree.Node{}
		ret = append(ret, make([]int, len(queue)))

		if i%2 == 0 {
			//从左到右
			for k, v := range queue {
				ret[i][k] = v.Value
			}
		} else {
			//从右到左
			for j := len(queue) - 1; j >= 0; j-- {
				ret[i][len(queue)-1-j] = queue[j].Value
			}
		}

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				//将孩子放入下层队列
				childQueue = append(childQueue, node.Left)
			}
			if node.Right != nil {
				childQueue = append(childQueue, node.Right)
			}
		}
	}
	return ret
}

/**
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

    3
   / \
  9  20
    /  \
   15   7


[
  [3],
  [9,20],
  [15,7]
]
*/
func LevelOrder(root *Tree.Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*Tree.Node{root}
	var ret [][]int
	var tmp []int
	for len(queue) > 0 {
		q := queue
		queue = nil
		tmp = []int{}
		for _, node := range q {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			tmp = append(tmp, node.Key)
		}
		ret = append(ret, tmp)
	}
	return ret
}
