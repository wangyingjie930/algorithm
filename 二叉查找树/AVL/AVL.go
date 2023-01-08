/**
  @author: wangyingjie
  @since: 2023/1/7
  @desc: 平衡二叉查找树
**/

package AVL

import (
	"fmt"
	"math"
)

type Node struct {
	Key   int
	Value int
	Left  *Node
	Right *Node
	Height int
}

func (node *Node) GetLeft() (interface{}, bool) {
	return node.Left, node.Left == nil
}

func (node *Node) GetRight() (interface{}, bool) {
	return node.Right, node.Right == nil
}

func (node *Node) GetKey() int {
	return node.Key
}

func (node *Node) GetValue() int {
	return node.Value
}

func NewNode(key int, value int) *Node {
	return &Node{Key: key, Value: value, Height: 0}
}

type Tree struct {
	root *Node
	count int
}

func (tree *Tree) Insert(key int, value int) {
	tree.root = insert(tree.root, key, value)
}

func insert(root *Node, key int, value int) *Node {
	if root == nil {
		return NewNode(key, value)
	}

	if root.Key > key {
		root.Left = insert(root.Left, key, value)
	}else if root.Key < key {
		root.Right = insert(root.Right, key, value)
	}else if key == root.Key {
		root.Value = value
	}

	if root.Left != nil && root.Right != nil {
		root.Height = int(math.Max(float64(root.Left.Height), float64(root.Right.Height)) + 1)
	}else if root.Left != nil {
		root.Height = root.Left.Height + 1
	}else if root.Right != nil {
		root.Height = root.Right.Height + 1
	}

	return root
}

/*
略显幼稚的层序遍历输出
func (tree *Tree) String() string {
	//queue: 相当于下一层准备好的队列
	queue := []*Node{tree.root}
	for level := 0; level < tree.root.Height; level++ {
		//每次进入一层之后, 就获取下一层的准备好的队列
		//q: 当前需要遍历的队列
		q := queue
		//并清空下一层队列, 准备在遍历q的过程中进行queue的收集, 为下一层的遍历做准备
		queue = nil

		fmt.Printf("\n第%d层\n", level)
		for len(q) > 0 {
			node := q[0]
			fmt.Print(strings.Repeat("\t", 13 - level*4), node)
			q = q[1:]
			if node.GetLeft != nil {
				//推入下一层
				queue = append(queue, node.GetLeft)
			}
			if node.GetRight != nil {
				queue = append(queue, node.GetRight)
			}
		}
	}
	return "\n\n"
}*/

/*更新节点高度*/
func (node *Node) updateHeight() int {
	if node == nil {
		return -1
	}
	node.Height = int(math.Max(float64(node.Left.updateHeight()), float64(node.Right.updateHeight())) + 1)
	return node.Height
}

func (node *Node) String() string {
	return fmt.Sprintf("节点: %d\t", node.Key)
}

// RightRotate 右旋
func (node *Node) RightRotate() *Node {
	temp := node.Left
	node.Left = temp.Right
	temp.Right = node
	temp.updateHeight()
	return temp
}

// LeftRotate 左旋
func (node *Node) LeftRotate() *Node {
	temp := node.Right
	node.Right = temp.Left
	temp.Left = node
	temp.updateHeight()
	return temp
}



