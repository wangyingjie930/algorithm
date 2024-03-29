/**
  @author: wangyingjie
  @since: 2023/2/16
  @desc
**/

package 岛屿数量

import "testing"

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{grid: [][]byte{
			{'0', '0', '1', '0', '0', '0', '0', '1', '0', '0', '0', '0', '0'},
			{'0', '0', '0', '0', '0', '0', '0', '1', '1', '1', '0', '0', '0'},
			{'0', '1', '1', '0', '1', '0', '0', '0', '0', '0', '0', '0', '0'},
			{'0', '1', '0', '0', '1', '1', '0', '0', '1', '0', '1', '0', '0'},
			{'0', '1', '0', '0', '1', '1', '0', '0', '1', '1', '1', '0', '0'},
			{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '0', '0'},
			{'0', '0', '0', '0', '0', '0', '0', '1', '1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0', '0', '0', '1', '1', '0', '0', '0', '0'},
		}}, want: 6},
		{name: "case1", args: args{grid: [][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.args.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
