/**
  @author: wangyingjie
  @since: 2023/2/17
  @desc
**/

package 验证二叉搜索树

import (
	"algorithm/二叉查找树/BST"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{nums: []int{1, 2, 4, 5, 6, 7}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := BST.NewBST()

			for _, v := range tt.args.nums {
				tree.Insert(v, v)
			}

			if got := isValidBST(tree.Root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
