/**
  @author: wangyingjie
  @since: 2023/2/15
  @desc: //TODO
**/

package 将有序数组转换为二叉搜索树

import (
	"fmt"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
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
			got := sortedArrayToBST(tt.args.nums)
			fmt.Print(got)
		})
	}
}
