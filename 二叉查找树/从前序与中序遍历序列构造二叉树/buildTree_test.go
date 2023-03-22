/**
  @author: wangyingjie
  @since: 2023/3/23
  @desc:
**/

package 从前序与中序遍历序列构造二叉树

import (
	"algorithm/二叉查找树/Tree"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{preorder: []int{3, 9, 20, 15, 7}, inorder: []int{9, 3, 15, 20, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := buildTree(tt.args.preorder, tt.args.inorder)
			Tree.PrintTree(node)
		})
	}
}
