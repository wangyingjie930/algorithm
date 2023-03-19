/**
  @author: wangyingjie
  @since: 2023/3/20
  @desc:
**/

package 二叉树的层序遍历

import (
	"algorithm/二叉查找树/Tree"
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "case1", args: args{nums: []int{1, Tree.NilNodeVal, 2, Tree.NilNodeVal, Tree.NilNodeVal, 3}}, want: [][]int{{1}, {2}, {3}}},
		{name: "case2", args: args{nums: []int{3, 9, 20, Tree.NilNodeVal, Tree.NilNodeVal, 15, 7}}, want: [][]int{{3}, {9, 20}, {15, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := Tree.NewTreeByNums(tt.args.nums, 0)
			Tree.PrintTree(node)

			if got := levelOrder(node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
