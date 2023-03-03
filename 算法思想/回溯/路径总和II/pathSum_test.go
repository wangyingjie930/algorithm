/**
  @author: wangyingjie
  @since: 2023/3/2
  @desc
**/

package 路径总和II

import (
	"algorithm/二叉查找树/Tree"
	"testing"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		nums      []int
		targetSum int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{nums: []int{
			5, 4, 8, 11, Tree.NilNodeVal, 13, 4, 7, 2, Tree.NilNodeVal, Tree.NilNodeVal,
			Tree.NilNodeVal, Tree.NilNodeVal, 5, 1,
		}, targetSum: 22}},
		{name: "case2", args: args{nums: []int{1, 2, 3}, targetSum: 5}},
		{name: "case3", args: args{nums: []int{-2, Tree.NilNodeVal, -3}, targetSum: -5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := Tree.NewTreeByNums(tt.args.nums, 0)
			Tree.PrintTree(node)
			got := pathSum(node, tt.args.targetSum)
			t.Log(got)
		})
	}
}
