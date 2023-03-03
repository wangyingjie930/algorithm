/**
  @author: wangyingjie
  @since: 2023/3/3
  @desc: //TODO
**/

package 使括号有效的最少添加

import "testing"

func Test_minAddToMakeValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{s: "())"}, want: 1},
		{name: "case2", args: args{s: "((("}, want: 3},
		{name: "case3", args: args{s: ")(()("}, want: 3},
		{name: "case4", args: args{s: "()))(("}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAddToMakeValid(tt.args.s); got != tt.want {
				t.Errorf("minAddToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
