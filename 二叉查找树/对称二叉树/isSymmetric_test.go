/**
  @author: wangyingjie
  @since: 2023/4/2
  @desc:
**/

package 对称二叉树

import (
	"algorithm/二叉查找树/Tree"
	"testing"
)

func Test_isSymmetric(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{nums: []int{1, 2, 2, 3, 4, 4, 3}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := Tree.NewTreeByNums(tt.args.nums, 0)
			Tree.PrintTree(node)

			if got := isSymmetric(node); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
