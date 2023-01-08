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
	// GetLeft nil和interface{}之间是不相等的, 必须显式地返回是否为nil
	GetLeft() (interface{}, bool)
	GetRight() (interface{}, bool)
	GetKey() int
	GetValue() int
}

// PrintTree /*打印二叉树*/
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
			q = append(q, entry{left.(TreeNode), r + 1, c - 1<<(height-r-1)})
		}
		if right, isnull := node.GetRight(); !isnull {
			q = append(q, entry{right.(TreeNode), r + 1, c + 1<<(height-r-1)})
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
				q = append(q, left.(TreeNode))
			}
			if right, isnull := node.GetRight(); !isnull {
				q = append(q, right.(TreeNode))
			}
		}
	}
	return h
}

