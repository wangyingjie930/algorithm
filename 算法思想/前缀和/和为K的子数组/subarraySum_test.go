/**
  @author: wangyingjie
  @since: 2023/4/2
  @desc:
**/

package 和为K的子数组

import "testing"

func Test_subarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{nums: []int{1, 1, 1}, k: 2}, want: 2},
		{name: "case1", args: args{nums: []int{1, 2, 3}, k: 3}, want: 2},
		{name: "case1", args: args{nums: []int{1, 2, 1, 2, 1}, k: 3}, want: 4},
		{name: "case1", args: args{nums: []int{1}, k: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
