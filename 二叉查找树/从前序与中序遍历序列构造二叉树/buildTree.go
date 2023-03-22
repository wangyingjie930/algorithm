/**
  @author: wangyingjie
  @since: 2023/3/23
  @desc: https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
  https://labuladong.github.io/algo/di-yi-zhan-da78c/shou-ba-sh-66994/dong-ge-da-172f0/
  掌握程度: 不熟
**/

package 从前序与中序遍历序列构造二叉树

import "algorithm/二叉查找树/Tree"

func buildTree(preorder []int, inorder []int) *Tree.Node {
	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func build(pre []int, preStart int, preEnd int, in []int, inStart int, inEnd int) *Tree.Node {
	if preStart > preEnd {
		return nil
	}

	//确定左右子树的分界线
	i := inStart
	for ; i <= inEnd; i++ {
		if in[i] == pre[preStart] {
			break
		}
	}

	leftSize := i - inStart //左子树个数

	//重点
	root := &Tree.Node{Value: pre[preStart], Key: pre[preStart]}
	//先序遍历: preStart+1 (左子树根节点位置), preStart+leftSize(整个左子树延续到哪个位置)
	//中序遍历: inStart (左子树根节点位置), i-1(整个左子树延续到哪个位置)
	root.Left = build(pre, preStart+1, preStart+leftSize, in, inStart, i-1)
	root.Right = build(pre, preStart+leftSize+1, preEnd, in, i+1, inEnd)
	return root
}
