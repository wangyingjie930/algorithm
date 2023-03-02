/**
  @author: wangyingjie
  @since: 2023/3/2
  @desc: //TODO
**/

package 数组中的第K个最大元素

import "testing"

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{nums: []int{3, 2, 1, 5, 6, 4}, k: 2}, want: 5},
		{name: "case2", args: args{nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, k: 4}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
