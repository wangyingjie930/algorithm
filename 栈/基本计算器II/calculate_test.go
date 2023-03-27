/**
  @author: wangyingjie
  @since: 2023/3/28
  @desc:
**/

package 基本计算器II

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{s: "1+2-3"}, want: 0},
		{name: "case1", args: args{s: "3+2*2"}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
