/**
  @author: wangyingjie
  @since: 2023/3/21
  @desc:
**/

package 二叉树中的最大路径和

import (
	"algorithm/二叉查找树/Tree"
	"testing"
)

func Test_maxPathSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{nums: []int{-2, -1}}, want: -1},
		{name: "case1", args: args{nums: []int{2, -1}}, want: 2},
		{name: "case1", args: args{nums: []int{1, 2}}, want: 3},
		{name: "case1", args: args{nums: []int{-3}}, want: -3},
		{name: "case1", args: args{nums: []int{1, -2, 3}}, want: 4},
		{name: "case1", args: args{nums: []int{-10, -2, -3, -1}}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := Tree.NewTreeByNums(tt.args.nums, 0)
			Tree.PrintTree(node)

			if got := maxPathSum(node); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
