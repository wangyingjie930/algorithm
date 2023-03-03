/**
  @author: wangyingjie
  @since: 2023/2/25
  @desc
**/

package 全排列

import (
	"fmt"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case", args: args{nums: []int{1, 2, 3}}},
		{name: "case", args: args{nums: []int{1, 2, 3, 4}}},
		{name: "case", args: args{nums: []int{0, 1}}},
		{name: "case", args: args{nums: []int{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.args.nums)
			fmt.Print(got)
		})
	}
}
