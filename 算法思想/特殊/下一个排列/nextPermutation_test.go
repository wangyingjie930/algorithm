/**
  @author: wangyingjie
  @since: 2023/3/3
  @desc:
**/

package 下一个排列

import "testing"

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{nums: []int{1, 2, 3}}},
		{name: "case2", args: args{nums: []int{3, 2, 1}}},
		{name: "case3", args: args{nums: []int{1, 1, 5}}},
		{name: "case4", args: args{nums: []int{1, 2}}},
		{name: "case5", args: args{nums: []int{1, 3, 2}}},
		{name: "case6", args: args{nums: []int{4, 3, 5, 7, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			t.Log(tt.args.nums)
		})
	}
}
