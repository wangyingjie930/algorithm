/**
  @author: wangyingjie
  @since: 2023/2/23
  @desc
**/

package 爬楼梯

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{n: 2}, want: 2},
		{name: "case1", args: args{n: 3}, want: 3},
		{name: "case1", args: args{n: 5}, want: 8},
		{name: "case1", args: args{n: 6}, want: 13},
		{name: "case7", args: args{n: 7}, want: 21},
		{name: "case8", args: args{n: 8}, want: 34},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairsDp(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbStairsDpSeven(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case6", args: args{n: 6}, want: 13},
		{name: "case7", args: args{n: 7}, want: 0},
		{name: "case8", args: args{n: 8}, want: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairsDpSeven(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbStairsBacktrack(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{n: 6}},
		{name: "case1", args: args{n: 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(climbStairsBacktrack(tt.args.n))
		})
	}
}
