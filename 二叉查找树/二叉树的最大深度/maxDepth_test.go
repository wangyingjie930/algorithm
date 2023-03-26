/**
  @author: wangyingjie
  @since: 2023/3/26
  @desc:
**/

package 二叉树的最大深度

import (
	"algorithm/二叉查找树/Tree"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{nums: []int{3, 9, 20, Tree.NilNodeVal, Tree.NilNodeVal, 15, 7}}, want: 3},
		{name: "case1", args: args{nums: []int{1}}, want: 1},
		{name: "case1", args: args{nums: []int{}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := Tree.NewTreeByNums(tt.args.nums, 0)
			if got := maxDepth(node); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
