/**
  @author: wangyingjie
  @since: 2023/3/28
  @desc: https://leetcode.cn/problems/maximum-width-of-binary-tree/description/
  掌握程度: 不熟
**/

package 二叉树最大宽度

import "algorithm/二叉查找树/Tree"

type treeNode struct {
	node   *Tree.Node
	number int
}

// widthOfBinaryTree
//  @Description: 核心, 广度遍历, 给节点加编号, 每一层开头和结尾的编号相减即获得最大宽度
//  @param root
//  @return int
func widthOfBinaryTree(root *Tree.Node) int {
	var queue []*treeNode
	tmp := []*treeNode{{node: root, number: 1}}
	maxLevel := 0

	for len(tmp) > 0 {
		//输出这一层的宽度
		if tmp[len(tmp)-1].number-tmp[0].number+1 > maxLevel {
			//获取最大的宽度
			maxLevel = tmp[len(tmp)-1].number - tmp[0].number + 1
		}

		queue = tmp
		tmp = []*treeNode{}

		for i := 0; i < len(queue); i++ {
			node := queue[i]
			if node.node.Left != nil {
				tmp = append(tmp, &treeNode{node: node.node.Left, number: 2 * node.number})
			}
			if node.node.Right != nil {
				tmp = append(tmp, &treeNode{node: node.node.Right, number: 2*node.number + 1})
			}
		}
	}

	return maxLevel
}
