/**
  @author: wangyingjie
  @since: 2023/3/2
  @desc: //TODO
**/

package 最长连续递增序列

import "testing"

func Test_findLengthOfLCIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{nums: []int{1, 3, 5, 4, 7}}, want: 3},
		{name: "case2", args: args{nums: []int{2, 2, 2, 2, 2}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLengthOfLCIS(tt.args.nums); got != tt.want {
				t.Errorf("findLengthOfLCIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
