/**
  @author: wangyingjie
  @since: 2023/2/18
  @desc: https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
**/

package 二叉树的最近公共祖先

import "algorithm/二叉查找树/Tree"

// lowestCommonAncestor
//  @Description: 核心点: 在递归的过程中只要遇到q/p就返回, 双方必定有交汇的地方, 当在某个节点判断左右两边不为空时,这个节点就是交汇的地方
//  @param root
//  @param p
//  @param q
//  @return *Tree.Node
func lowestCommonAncestor(root, p, q *Tree.Node) *Tree.Node {
	if root == p || root == q {
		//只要遇到q/p就返回
		return root
	}
	if root == nil {
		return nil
	}

	//子问题: root左边是否有p/q
	leftNode := lowestCommonAncestor(root.Left, p, q)
	//右边是否有p/q
	rightNode := lowestCommonAncestor(root.Right, p, q)
	if (leftNode == p && rightNode == q) || (leftNode == q && rightNode == p) {
		//两边都有的话, 返回root
		return root
	} else if leftNode != nil {
		return leftNode
	} else if rightNode != nil {
		return rightNode
	} else {
		return nil
	}
}

/**
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
*/
