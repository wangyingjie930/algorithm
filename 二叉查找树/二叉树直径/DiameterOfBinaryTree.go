/**
  @author: wangyingjie
  @since: 2023/2/16
  @desc: https://leetcode.cn/problems/diameter-of-binary-tree/description/
**/

package 二叉树直径

import (
	"algorithm/二叉查找树/Tree"
)

func DiameterOfBinaryTree(root *Tree.Node) int {
	//1. 当x参与的时候: 取 左的子树高度 + 右子树的高度 + 1
	//2. 当x不参与的时候: 取 左子树最长距离 或者 右子树最长距离
	//3种取最大值
	_, distance := processDiameter(root)
	return distance
}

/**
树形dp套路第一步：
以某个节点X为头节点的子树中，分析答案有哪些可能性，并且这种分析是以X的左子树、X的右子树和X整棵树的角度来考虑可能性的
树形dp套路第二步：
根据第一步的可能性分析，列出所有需要的信息
树形dp套路第三步：
合并第二步的信息，对左树和右树提出同样的要求，并写出信息结构
树形dp套路第四步：
设计递归函数，递归函数是处理以X为头节点的情况下的答案。
包括设计递归的basecase，默认直接得到左树和右树的所有信息，以及把可能性做整合，并且要返回第三步的信息结构这四个小步骤
*/
func processDiameter(node *Tree.Node) (int, int) {
	if node == nil {
		return 0, 0
	}
	leftHeight, leftDistance := processDiameter(node.Left)
	rightHeight, rightDistance := processDiameter(node.Right)

	maxDistance := leftHeight + rightHeight + 1
	maxHeight := leftHeight
	if maxDistance < leftDistance {
		maxDistance = leftDistance
	}
	if maxDistance < rightDistance {
		maxDistance = rightDistance
	}
	if maxHeight < rightHeight {
		maxHeight = rightHeight
	}
	return maxHeight + 1, maxDistance
}
