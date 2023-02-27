/**
  @author: wangyingjie
  @since: 2023/2/27
  @desc: //TODO
**/

package 移动零

import "testing"

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{nums: []int{0, 1, 0, 3, 12}}},
		{name: "case2", args: args{nums: []int{0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			t.Log(tt.args.nums)
		})
	}
}
