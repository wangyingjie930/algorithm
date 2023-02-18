/**
  @author: wangyingjie
  @since: 2022/11/27
  @desc: //TODO
**/

package Morris

import (
	"algorithm/二叉查找树/Tree"
	"testing"
)

func Test_morrisIn(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{nums: []int{1, 2, 4, 5, 6, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := Tree.NewTreeByNums(tt.args.nums, 0)
			Tree.PrintTree(node)

			morrisIn(node)
		})
	}
}
