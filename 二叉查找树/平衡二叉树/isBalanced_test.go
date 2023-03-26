/**
  @author: wangyingjie
  @since: 2023/3/26
  @desc:
**/

package 平衡二叉树

import (
	"algorithm/二叉查找树/Tree"
	"testing"
)

func Test_isBalanced(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{nums: []int{3, 9, 20, Tree.NilNodeVal, Tree.NilNodeVal, 15, 7}}, want: false},
		{name: "case1", args: args{nums: []int{1}}, want: true},
		{name: "case1", args: args{nums: []int{}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := Tree.NewTreeByNums(tt.args.nums, 0)
			if got := isBalanced(node); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
