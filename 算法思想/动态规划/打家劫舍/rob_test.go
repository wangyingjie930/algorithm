/**
  @author: wangyingjie
  @since: 2023/4/2
  @desc:
**/

package 打家劫舍

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{nums: []int{1, 2, 3, 1}}, want: 4},
		{name: "case1", args: args{nums: []int{2, 7, 9, 3, 1}}, want: 12},
		{name: "case1", args: args{nums: []int{2, 1, 1, 2}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
