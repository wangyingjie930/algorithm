/**
  @author: wangyingjie
  @since: 2023/3/3
  @desc: //TODO
**/

package 二叉树的右视图

import (
	"algorithm/二叉查找树/Tree"
	"reflect"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "case1", args: args{nums: []int{1, 2, 3, Tree.NilNodeVal, 5, Tree.NilNodeVal, 4}}, want: []int{1, 3, 4}},
		{name: "case2", args: args{nums: []int{1, Tree.NilNodeVal, 3}}, want: []int{1, 3}},
		{name: "case3", args: args{nums: []int{}}, want: []int{}},
		{name: "case4", args: args{nums: []int{1, 2}}, want: []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := Tree.NewTreeByNums(tt.args.nums, 0)
			if got := rightSideView(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
