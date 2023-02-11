/**
  @author: wangyingjie
  @since: 2023/2/10
  @desc: //TODO
**/

package 颜色分类

import (
	"fmt"
	"testing"
)

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{nums: []int{2, 0, 2, 1, 1, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
			fmt.Print(tt.args.nums)
		})
	}
}
