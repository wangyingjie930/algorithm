/**
  @author: wangyingjie
  @since: 2023/3/26
  @desc:
**/

package 求根节点到叶节点数字之和

import (
	"algorithm/二叉查找树/Tree"
	"testing"
)

func Test_sumNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{nums: []int{1, 2, 3}}, want: 25},
		{name: "case1", args: args{nums: []int{4, 9, 0, 5, 1}}, want: 1026},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := Tree.NewTreeByNums(tt.args.nums, 0)
			if got := sumNumbers(node); got != tt.want {
				t.Errorf("sumNumbers() = %v, want %v", got, tt.want)
			}
			t.Log(bfs(node))
		})
	}
}
