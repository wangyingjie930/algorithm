package BST

import "algorithm/二叉查找树/Tree"

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
