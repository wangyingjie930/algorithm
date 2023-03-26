/**
  @author: wangyingjie
  @since: 2023/3/26
  @desc:
**/

package 子集

import (
	"testing"
)

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{nums: []int{1, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(subsets(tt.args.nums))
		})
	}
}
