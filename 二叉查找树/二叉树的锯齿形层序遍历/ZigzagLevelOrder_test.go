/**
  @author: wangyingjie
  @since: 2023/2/18
**/

package 二叉树的锯齿形层序遍历

import (
	"algorithm/二叉查找树/Tree"
	"fmt"
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{nums: []int{3, 9, 20, Tree.NilNodeVal, Tree.NilNodeVal, 15, 7}}},
		{name: "case1", args: args{nums: []int{1}}},
		{name: "case1", args: args{nums: []int{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := Tree.NewTreeByNums(tt.args.nums, 0)
			fmt.Print(zigzagLevelOrder(node), "\n")
		})
	}
}
