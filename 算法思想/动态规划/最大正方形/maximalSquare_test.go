/**
  @author: wangyingjie
  @since: 2023/3/26
  @desc:
**/

package 最大正方形

import "testing"

func Test_maximalSquare(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{[][]byte{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		}}, want: 4},
		{name: "case1", args: args{[][]byte{
			{'0', '1'},
			{'1', '0'},
		}}, want: 1},
		{name: "case1", args: args{[][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '1', '1', '0'},
			{'1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'0', '0', '1', '1', '1'},
		}}, want: 16},

		{name: "case1", args: args{[][]byte{
			{'0', '0', '0', '1'},
			{'1', '1', '0', '1'},
			{'1', '1', '1', '1'},
			{'0', '1', '1', '1'},
			{'0', '1', '1', '1'},
		}}, want: 9},

		{name: "case1", args: args{[][]byte{
			{'1', '1', '1', '1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1', '1', '1', '0'},
			{'1', '1', '1', '1', '1', '1', '1', '0'},
			{'1', '1', '1', '1', '1', '0', '0', '0'},
			{'0', '1', '1', '1', '1', '0', '0', '0'},
		}}, want: 16},

		{name: "case1", args: args{[][]byte{
			{'0', '1', '1', '0', '1'},
			{'1', '1', '0', '1', '0'},
			{'0', '1', '1', '1', '0'},
			{'1', '1', '1', '1', '0'},
			{'1', '1', '1', '1', '1'},
			{'0', '0', '0', '0', '0'},
		}}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquare(tt.args.matrix); got != tt.want {
				t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
