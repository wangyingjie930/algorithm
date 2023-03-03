/**
  @author: wangyingjie
  @since: 2023/3/2
  @desc
**/

package 三数之和

import (
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{nums: []int{-1, 0, 1, 2, -1, -4}}},
		{name: "case2", args: args{nums: []int{0, 1, 1}}},
		{name: "case3", args: args{nums: []int{-2, 0, 0, 2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.args.nums)
			t.Log(got)
		})
	}
}
