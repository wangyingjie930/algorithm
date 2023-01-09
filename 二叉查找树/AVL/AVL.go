/**
  @author: wangyingjie
  @since: 2023/1/7
  @desc: 平衡二叉查找树
**/

package AVL

import (
	BST "algorithm/二叉查找树"
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

func (node *Node) GetLeft() (BST.TreeNode, bool) {
	return node.Left, node.Left == nil
}

func (node *Node) GetRight() (BST.TreeNode, bool) {
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

//插入的时候调整平衡性
func insert(root *Node, key int, value int) *Node {
	if root == nil {
		return NewNode(key, value)
	}

	//选取对应的位置插入
	if root.Key > key {
		root.Left = insert(root.Left, key, value)
	}else if root.Key < key {
		root.Right = insert(root.Right, key, value)
	}else if key == root.Key {
		root.Value = value
	}

	//进行高度的更新
	if root.Left != nil && root.Right != nil {
		root.Height = int(math.Max(float64(root.Left.Height), float64(root.Right.Height)) + 1)
	}else if root.Left != nil {
		root.Height = root.Left.Height + 1
	}else if root.Right != nil {
		root.Height = root.Right.Height + 1
	}

	//进行平衡性的维护
	root = root.maintain()

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

// 进行平衡性的维护
func (node *Node) maintain() *Node {
	leftHeight := 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}
	rightHeight := 0
	if node.Right != nil {
		rightHeight = node.Right.Height
	}
	if math.Abs(float64(rightHeight)-float64(leftHeight)) <= 1 {
		//平衡因子小于等于1则表示平衡
		return node
	}

	if leftHeight > rightHeight {
		//左树高, 那么最终肯定会右旋
		leftNode := node.Left
		leftLeftHeight, leftRightHeight := 0, 0
		if leftNode.Left != nil {
			leftLeftHeight = leftNode.Left.Height
		}
		if leftNode.Right != nil {
			leftRightHeight = leftNode.Right.Height
		}
		//如果左节点(也就是待上升的节点的)的右子树比较高, 单纯旋转的话, 待下降的节点会继承待上升节点的右节点, 高度不会降下来, 需要先左后右
		if leftRightHeight > leftLeftHeight {
			node.Left = leftNode.LeftRotate()
		}
		//如果左节点(也就是待上升的节点的)的左子树比较高, 可以直接一次右旋就可以完成
		//这里必须要返回右旋之后的根节点
		return node.RightRotate()
	}else if rightHeight > leftHeight {
		//右树高, 那么最终肯定会左旋
		rightNode := node.Right
		rightLeftHeight, rightRightHeight := 0, 0
		if rightNode.Left != nil {
			rightLeftHeight = rightNode.Left.Height
		}
		if rightNode.Right != nil {
			rightRightHeight = rightNode.Right.Height
		}
		if rightLeftHeight > rightRightHeight {
			node.Right = rightNode.RightRotate()
		}
		return node.LeftRotate()
	}
	return node
}



