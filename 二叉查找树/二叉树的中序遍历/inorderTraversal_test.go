/**
  @author: wangyingjie
  @since: 2023/3/20
  @desc:
**/

package 二叉树的中序遍历

import (
	"algorithm/二叉查找树/Tree"
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "case1", args: args{nums: []int{1, Tree.NilNodeVal, 2, Tree.NilNodeVal, Tree.NilNodeVal, 3}}, want: []int{1, 3, 2}},
		{name: "case2", args: args{nums: []int{1}}, want: []int{1}},
	}
	for _, tt := range tests {
		// 三种方法进行测试
		for _, v := range []inorder{new(Recur), new(Iteration), new(morris)} {
			t.Run(tt.name, func(t *testing.T) {
				node := Tree.NewTreeByNums(tt.args.nums, 0)
				Tree.PrintTree(node)

				t.Log(reflect.TypeOf(v))
				if got := v.inorderTraversal(node); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
