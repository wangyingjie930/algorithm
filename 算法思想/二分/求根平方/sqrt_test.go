/**
  @author: wangyingjie
  @since: 2023/2/22
  @desc
**/

package 求根平方

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{x: 16}, want: 4},
		{name: "case1", args: args{x: 20164}, want: 142},
		{name: "case1", args: args{x: 981130329}, want: 31323},
		{name: "case1", args: args{x: 10001}, want: 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mySqrtAccuracy(t *testing.T) {
	type args struct {
		x int
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "case1", args: args{x: 16, n: 2}, want: 4},
		{name: "case1", args: args{x: 17, n: 3}, want: 4.123},
		{name: "case1", args: args{x: 18, n: 4}, want: 4.2426},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrtAccuracy(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("mySqrtAccuracy() = %v, want %v", got, tt.want)
			}
		})
	}
}
