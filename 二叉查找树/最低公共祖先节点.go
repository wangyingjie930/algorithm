// Package BST
// @Description: 给定两个二叉树的节点node1和node2，找到他们的最低公共祖先节点
package BST

import "algorithm/二叉查找树/Tree"

func LowestAncestor(head, o1, o2 *Tree.Node) *Tree.Node {
	//当当前节点为空返回
	//当当前节点为要搜索的其中一个节点时返回, 表明在当前的子树中包含该节点, 并返回该节点, 层层往上
	//当遇到LCA时, 这两个往上的节点相遇, 则识别出来LCA
	if head == nil || head == o1 || head == o2 {
		return head
	}

	left := LowestAncestor(head.Left, o1, o2)
	right := LowestAncestor(head.Right, o1, o2)

	if left != nil && right != nil {
		//如果一边包含o1, 一边包含o2, 那么这个节点就是最低公共祖先节点
		return head
	} else if left == nil {
		//左边没有包含节点, 右边有则返回右边的(o1/o2)
		return right
	}
	//右边没有包含节点, 左边有则返回左边的(o1/o2)
	return left
}
