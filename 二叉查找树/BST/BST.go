package BST

import (
	"algorithm/二叉查找树/Tree"
	"fmt"
)

type BST struct {
	Root  *Tree.Node
	count int
}

func NewBST() BST {
	return BST{Root: nil, count: 0}
}

func (b *BST) IsEmpty() bool {
	return b.count == 0
}

func (b *BST) Size() int {
	return b.count
}

func (b *BST) Insert(key, value int) {
	b.Root = b.insert(b.Root, key, value)
}

func (b *BST) Contain(key int) bool {
	return b.contain(b.Root, key)
}

func (b *BST) Search(key int) *int {
	return b.search(b.Root, key)
}

//前序遍历
func (b *BST) PreOrder() {
	preOrder(b.Root)
}

//中序遍历
func (b *BST) InOrder() {
	inOrder(b.Root)
}

func (b *BST) PostOrder() {
	postOrder(b.Root)
}

func (b *BST) LevelOrder() {

}

func (b *BST) Minimum() int {
	node := minimum(b.Root)
	return node.Key
}

func (b *BST) Maximum() int {
	node := maximum(b.Root)
	return node.Key
}

func (b *BST) RemoveMin() {
	if b.Root != nil {
		removeMin(b.Root)
	}
}

func (b *BST) RemoveMax() {
	if b.Root != nil {
		removeMax(b.Root)
	}
}

func (b *BST) Remove(key int) {
	b.Root = remove(b.Root, key)
}

//思考: 二分搜索树的floor和ceil (最小大于58的节点和最大小于58的节点是什么)
//思考: 二分搜索树的rank(58在树中排第几)和select(第10名的节点是什么), 对每个node添加属性: 左右子树加上自身的数量, 保存一份
//思考: 支持重复元素的树, node添加属性: 重复个数

/**
插入节点
*/
func (b *BST) insert(node *Tree.Node, key, value int) *Tree.Node {
	if node == nil {
		b.count++
		return &Tree.Node{Key: key, Value: value, Right: nil, Left: nil}
	}

	if node.Key == key {
		node.Value = value
	} else if node.Key > key {
		node.Left = b.insert(node.Left, key, value)
	} else {
		node.Right = b.insert(node.Right, key, value)
	}
	return node
}

/**
key是否包含在树中
*/
func (b *BST) contain(node *Tree.Node, key int) bool {
	if node == nil {
		return false
	}
	if node.Key == key {
		return true
	} else if node.Key > key {
		return b.contain(node.Left, key)
	} else {
		return b.contain(node.Right, key)
	}
}

/**
返回查找节点的值
*/
func (b *BST) search(node *Tree.Node, key int) *int {
	if node == nil {
		return nil
	}
	if node.Key == key {
		return &node.Value
	} else if node.Key > key {
		return b.search(node.Left, key)
	} else {
		return b.search(node.Right, key)
	}
}

/**
前序遍历
*/
func preOrder(node *Tree.Node) {
	if node == nil {
		return
	}
	fmt.Print(node.Key, " ")

	preOrder(node.Left)
	preOrder(node.Right)
}

/**
中序遍历
*/
func inOrder(node *Tree.Node) {
	if node == nil {
		return
	}
	inOrder(node.Left)
	fmt.Print(node.Key, " ")
	inOrder(node.Right)
}

/**
后序遍历
*/
func postOrder(node *Tree.Node) {
	if node == nil {
		return
	}
	postOrder(node.Left)
	postOrder(node.Right)
	fmt.Print(node.Key, " ")
}

/**
最小节点, 递归获取左节点就可以找到
*/
func minimum(node *Tree.Node) *Tree.Node {
	if node.Left == nil {
		return node
	}
	return minimum(node.Left)
}

/**
最大节点, 递归获取右节点可以找到
*/
func maximum(node *Tree.Node) *Tree.Node {
	if node.Right == nil {
		return node
	}
	return maximum(node.Right)
}

/**
删除最大节点
规律: 递归获取右节点可以找到, 而且它最多也只有左子树
*/
func removeMax(node *Tree.Node) *Tree.Node {
	if node.Right == nil {
		left := node.Left
		node = nil
		return left
	}
	node.Right = removeMax(node.Right)
	return node
}

/**
删除最小节点
规律: 递归获取左节点就可以找到, 而且它做多只有右子树
*/
func removeMin(node *Tree.Node) *Tree.Node {
	if node.Left == nil {
		right := node.Right
		node = nil
		return right
	}
	node.Left = removeMin(node.Left)
	return node
}

/**
删除节点
*/
func remove(node *Tree.Node, key int) *Tree.Node {
	if node == nil {
		return nil
	}
	if key < node.Key {
		//递归向左找
		node.Left = remove(node.Left, key)
		return node
	}else if key > node.Key {
		//递归向右找
		node.Right = remove(node.Right, key)
		return node
	}else {
		//查找到节点
		if node.Right == nil {
			//只含有左子树
			leftNode := node.Left
			node = nil
			return leftNode
		}else if node.Left == nil {
			//只含有右子树
			rightNode := node.Right
			node = nil
			return rightNode
		}else {
			//同时又左右子树
			//找出后继节点作为新的节点(被删除节点的右子树的最小节点)
			successor := minimum(node.Right)
			//重新赋值新节点的左右节点(被删除节点的左右节点)
			successor.Right = removeMin(node.Right)
			successor.Left = node.Left
			node = nil
			return successor
		}
	}
}
