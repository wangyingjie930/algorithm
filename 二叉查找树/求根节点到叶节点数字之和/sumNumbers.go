/**
  @author: wangyingjie
  @since: 2023/3/26
  @desc: https://leetcode.cn/problems/sum-root-to-leaf-numbers/
**/

package 求根节点到叶节点数字之和

import "algorithm/二叉查找树/Tree"

func sumNumbers(root *Tree.Node) int {
	sum := 0
	dfs(root, 0, &sum)
	return sum
}

// dfs
//  @Description: 解法1, 深度遍历
//  @param root
//  @param num
//  @param sum
func dfs(root *Tree.Node, num int, sum *int) {
	if root.Left == nil && root.Right == nil {
		*sum = *sum + (num*10 + root.Value)
		return
	}

	num = num*10 + root.Value
	if root.Left != nil {
		dfs(root.Left, num, sum)
	}
	if root.Right != nil {
		dfs(root.Right, num, sum)
	}
}

//  pair
//  @Description: 解法二: 广度优先
type pair struct {
	node *Tree.Node
	num  int
}

func bfs(root *Tree.Node) int {
	var queue []*pair
	nextQueue := []*pair{{node: root, num: root.Value}}
	sum := 0

	for len(nextQueue) > 0 {
		queue = nextQueue
		nextQueue = []*pair{}
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.node.Left != nil {
				nextQueue = append(nextQueue, &pair{num: node.num*10 + node.node.Left.Value, node: node.node.Left})
			}
			if node.node.Right != nil {
				nextQueue = append(nextQueue, &pair{num: node.num*10 + node.node.Right.Value, node: node.node.Right})
			}
			if node.node.Left == nil && node.node.Right == nil {
				sum += node.num
			}
		}
	}
	return sum
}
