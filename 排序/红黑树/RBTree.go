package BST

import BST "imooc-product/backend/二叉查找树"

const RED bool = false
const BLEAK bool = true

type RBNode struct {
	BST.Node
	color bool
}

type RBTree struct {
	BST.BST
}

func (rb *RBTree)Color(node *RBNode, color bool) *RBNode  {
	if node == nil {
		return node
	}
	node.color = color
	return node
}