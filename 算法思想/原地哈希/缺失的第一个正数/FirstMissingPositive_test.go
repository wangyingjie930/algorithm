/**
  @author: wangyingjie
  @since: 2023/2/16
  @desc
**/

package 缺失的第一个正数

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case2", args: args{nums: []int{1, 2, 0}}, want: 3},
		{name: "case2", args: args{nums: []int{3, 4, -1, 1}}, want: 2},
		{name: "case2", args: args{nums: []int{7, 8, 9, 11, 12}}, want: 1},
		{name: "case2", args: args{nums: []int{-1, 4, 2, 1, 9, 10}}, want: 3},
		{name: "case2", args: args{nums: []int{1}}, want: 2},
		{name: "case2", args: args{nums: []int{0, 1, 2}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
