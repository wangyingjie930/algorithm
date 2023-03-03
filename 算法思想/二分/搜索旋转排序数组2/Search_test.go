/**
  @author: wangyingjie
  @since: 2023/2/9
  @desc
**/

package 搜索旋转排序数组2

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{nums: []int{3, 4, 5, 1, 2}, target: 1}, want: true},
		{name: "case2", args: args{nums: []int{1, 0, 1, 1, 1}, target: 0}, want: true},
		{name: "case2", args: args{nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 2}, target: 2}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
