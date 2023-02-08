/**
  @author: wangyingjie
  @since: 2023/2/8
  @desc: //TODO
**/

package 最大子数组和

import "testing"

func Test_maxSubArrayDynamic(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArrayDynamic(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArrayDynamic() = %v, want %v", got, tt.want)
			}
		})
	}
}
