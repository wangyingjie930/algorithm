/**
  @author: wangyingjie
  @since: 2023/2/16
  @desc
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
		{name: "case1", args: args{nums: []int{1, 2, 3, 4, 5}}, want: 3},
		{name: "case1", args: args{nums: []int{1, 2}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := Tree.NewTreeByNums(tt.args.nums, 0)
			Tree.PrintTree(node)

			if got := diameterOfBinaryTree(node); got != tt.want {
				t.Errorf("DiameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
