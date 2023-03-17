/**
  @author: wangyingjie
  @since: 2023/3/17
  @desc:
**/

package 手撕快排

import (
	"testing"
)

func Test_sortArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortArray(tt.args.nums)
			t.Log(got)
		})
	}
}
