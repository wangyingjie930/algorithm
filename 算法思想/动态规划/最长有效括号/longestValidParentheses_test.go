/**
  @author: wangyingjie
  @since: 2023/3/3
  @desc
**/

package 最长有效括号

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{s: "(()"}, want: 2},
		{name: "case2", args: args{s: ")()())"}, want: 4},
		{name: "case3", args: args{s: ""}, want: 0},
		{name: "case4", args: args{s: "()(()"}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
