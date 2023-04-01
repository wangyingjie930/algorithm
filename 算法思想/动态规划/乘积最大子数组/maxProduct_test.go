/**
  @author: wangyingjie
  @since: 2023/4/2
  @desc:
**/

package 乘积最大子数组

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{nums: []int{2, 3, -2, 4}}, want: 6},
		{name: "case1", args: args{nums: []int{-2, 1, -1}}, want: 2},
		{name: "case1", args: args{nums: []int{-2, 0, -1}}, want: 0},
		{name: "case1", args: args{nums: []int{-2}}, want: -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
