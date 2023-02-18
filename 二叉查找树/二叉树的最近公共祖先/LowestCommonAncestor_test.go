/**
  @author: wangyingjie
  @since: 2023/2/18
  @desc: //TODO
**/

package 二叉树的最近公共祖先

import (
	"algorithm/二叉查找树/Tree"
	"fmt"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		nums []int
		p    int
		q    int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{nums: []int{3, 5, 1, 6, 2, 0, 8, Tree.NilNodeVal, Tree.NilNodeVal, 7, 4}, p: 5, q: 1}},
		{name: "case1", args: args{nums: []int{3, 5, 1, 6, 2, 0, 8, Tree.NilNodeVal, Tree.NilNodeVal, 7, 4}, p: 5, q: 4}},
		{name: "case1", args: args{nums: []int{1, 2}, p: 1, q: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := Tree.NewTreeByNums(tt.args.nums, 0)
			pNode := findNode(root, tt.args.p)
			qNode := findNode(root, tt.args.q)
			got := lowestCommonAncestor(root, pNode, qNode)
			fmt.Print(got.Value, "\n")
			Tree.PrintTree(root)
		})
	}
}

// findNode
//  @Description: 定位节点位置
//  @param root
//  @param value
//  @return *Tree.Node
func findNode(root *Tree.Node, value int) *Tree.Node {
	if root == nil {
		return nil
	}
	if root.Value == value {
		return root
	}

	left := findNode(root.Left, value)
	right := findNode(root.Right, value)

	if left != nil {
		return left
	} else if right != nil {
		return right
	} else {
		return nil
	}
}
