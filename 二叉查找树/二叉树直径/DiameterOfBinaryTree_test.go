/**
  @author: wangyingjie
  @since: 2023/2/16
  @desc: //TODO
**/

package 二叉树直径

import (
	"algorithm/二叉查找树/Tree"
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{nums: []int{1, 2, 4, 5, 6, 7}}, want: 5},
		{name: "case1", args: args{nums: []int{4, 2, 5, 6, 7, 1}}, want: 5},
		{name: "case1", args: args{nums: []int{1, 5, 9, 6, 7, Tree.NilNodeVal, Tree.NilNodeVal,
			Tree.NilNodeVal, Tree.NilNodeVal, Tree.NilNodeVal, 8}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := Tree.NewTreeByNums(tt.args.nums, 0)
			Tree.PrintTree(node)

			if got := DiameterOfBinaryTree(node); got != tt.want {
				t.Errorf("DiameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
