/**
  @author: wangyingjie
  @since: 2023/2/18
  @desc: //TODO
**/

package 搜索旋转排序数组

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0}, want: 4},
		{name: "case1", args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3}, want: -1},
		{name: "case1", args: args{nums: []int{1}, target: 0}, want: -1},
		{name: "case1", args: args{nums: []int{1, 3, 5}, target: 3}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
