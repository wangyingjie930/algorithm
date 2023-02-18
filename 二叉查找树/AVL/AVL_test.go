/**
  @author: wangyingjie
  @since: 2023/1/7
  @desc: //TODO
**/

package AVL

import (
	BST "algorithm/二叉查找树/Tree"
	"testing"
)

var tree = newAVL()

func newAVL() *Tree {
	tree := Tree{root: nil, count: 0}
	arr := []int{385, 823, 441, 805, 258, 765, 868, 83, 11, 22}
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for k, v := range arr {
		tree.Insert(arr1[k], v)
	}

	return &tree
}

func TestNode_leftRotate(t *testing.T) {
	BST.PrintTree(tree.root)
	tree.root = tree.root.RightRotate()
	BST.PrintTree(tree.root)
}
