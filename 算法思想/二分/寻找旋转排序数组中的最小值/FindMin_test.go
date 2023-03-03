/**
  @author: wangyingjie
  @since: 2023/2/8
  @desc
**/

package 寻找旋转排序数组中的最小值

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{nums: []int{3, 4, 5, 1, 2}}, want: 1},
		{name: "case1", args: args{nums: []int{0, 1, 2, 4, 5, 6, 7}}, want: 0},
		{name: "case1", args: args{nums: []int{3, 1, 2}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
