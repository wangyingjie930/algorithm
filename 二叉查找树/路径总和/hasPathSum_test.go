/**
  @author: wangyingjie
  @since: 2023/4/2
  @desc:
**/

package 路径总和

import (
	"algorithm/二叉查找树/Tree"
	"testing"
)

func Test_hasPathSum(t *testing.T) {
	type args struct {
		nums      []int
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{nums: []int{1, 2, 3}, targetSum: 5}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := Tree.NewTreeByNums(tt.args.nums, 0)
			Tree.PrintTree(node)

			if got := hasPathSum(node, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
