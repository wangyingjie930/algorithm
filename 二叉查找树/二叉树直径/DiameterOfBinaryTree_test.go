/**
  @author: wangyingjie
  @since: 2023/2/16
  @desc: //TODO
**/

package 二叉树直径

import (
	BST "algorithm/二叉查找树"
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
		{name: "case1", args: args{nums: []int{1, 2, 4, 5, 6, 7}}, want: 6},
		{name: "case1", args: args{nums: []int{4, 2, 5, 6, 7, 1}}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := BST.NewBST()
			for _, v := range tt.args.nums {
				tree.Insert(v, v)
			}
			if got := DiameterOfBinaryTree(tree.Root); got != tt.want {
				t.Errorf("DiameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
