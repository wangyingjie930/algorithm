/**
  @author: wangyingjie
  @since: 2023/1/8
  @desc: //TODO
**/

package BST

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode interface {
	// GetLeft 坑点: nil和interface{}之间是不相等的, 必须显式地返回是否为nil
	// 返回的值也只要TreeNode, 不需要*TreeNode!!
	GetLeft() (TreeNode, bool)
	GetRight() (TreeNode, bool)
	GetKey() int
	GetValue() int
}

// PrintTree /*打印二叉树*/
// 坑点: 接收的参数为TreeNode, 而非*TreeNode!!
// 这里的TreeNode传值的话, 那么实现接口的结构体接收者要以值的形式实现接口
// 如果传指针的话, 那么实现接口的结构体接收者要以指针的形式实现接口
// 传interface，传递的是一个interface对象，这个对象占用16字节长度，包含一个指向原数据的指针，和一个指向运行时类型信息的指针
func PrintTree(root TreeNode)  {
	height := calDepth(root)
	m := height + 1
	n := 1<<m - 1
	ans := make([][]string, m)
	for i := range ans {
		ans[i] = make([]string, n)
	}
	type entry struct {
		node TreeNode
		r, c int
	}
	q := []entry{{root, 0, (n - 1) / 2}}
	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		node, r, c := e.node, e.r, e.c
		ans[r][c] = strconv.Itoa(node.GetKey())
		if left, isnull := node.GetLeft(); !isnull {
			q = append(q, entry{left, r + 1, c - 1<<(height-r-1)})
		}
		if right, isnull := node.GetRight(); !isnull {
			q = append(q, entry{right, r + 1, c + 1<<(height-r-1)})
		}
	}

	for _, levels := range ans {
		fmt.Println(strings.Join(levels[:], "-"))
	}
	fmt.Println("")
}

/*水平遍历获取某个节点的高度*/
func calDepth(root TreeNode) int {
	h := -1
	q := []TreeNode{root}
	for len(q) > 0 {
		h++
		tmp := q
		q = nil
		for _, node := range tmp {
			if left, isnull := node.GetLeft(); !isnull {
				q = append(q, left)
			}
			if right, isnull := node.GetRight(); !isnull {
				q = append(q, right)
			}
		}
	}
	return h
}

