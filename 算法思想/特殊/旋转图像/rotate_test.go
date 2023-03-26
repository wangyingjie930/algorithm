/**
  @author: wangyingjie
  @since: 2023/3/26
  @desc:
**/

package 旋转图像

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{matrix: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("before: ", tt.args.matrix)
			rotate(tt.args.matrix)
			t.Log("after: ", tt.args.matrix)
		})
	}
}
