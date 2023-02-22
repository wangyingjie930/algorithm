/**
  @author: wangyingjie
  @since: 2023/2/16
  @desc: //TODO
**/

package 搜索二维矩阵

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{matrix: [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}, target: 3}, want: true},

		{name: "case1", args: args{matrix: [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}, target: 13}, want: false},

		{name: "case1", args: args{matrix: [][]int{
			{1}, {3},
		}, target: 1}, want: true},

		{name: "case1", args: args{matrix: [][]int{
			{1}, {3}, {5},
		}, target: 3}, want: true},

		{name: "case1", args: args{matrix: [][]int{
			{1},
		}, target: 0}, want: false},
		//[[1,3,5,7],[10,11,16,20],[23,30,34,50]]
		{name: "case1", args: args{matrix: [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 50},
		}, target: 11}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
