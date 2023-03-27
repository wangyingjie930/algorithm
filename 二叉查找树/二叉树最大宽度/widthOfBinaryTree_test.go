/**
  @author: wangyingjie
  @since: 2023/3/28
  @desc:
**/

package 二叉树最大宽度

import (
	"algorithm/二叉查找树/Tree"
	"testing"
)

func Test_widthOfBinaryTree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case2", args: args{nums: []int{1, 3, 2, 5, 3, Tree.NilNodeVal, 9}}, want: 4},
		{name: "case2", args: args{nums: []int{
			1, 3, 2, 5, Tree.NilNodeVal, Tree.NilNodeVal, 9, 6,
			Tree.NilNodeVal, Tree.NilNodeVal, Tree.NilNodeVal, Tree.NilNodeVal, Tree.NilNodeVal, 7,
		}}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := Tree.NewTreeByNums(tt.args.nums, 0)
			Tree.PrintTree(node)

			if got := widthOfBinaryTree(node); got != tt.want {
				t.Errorf("widthOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
