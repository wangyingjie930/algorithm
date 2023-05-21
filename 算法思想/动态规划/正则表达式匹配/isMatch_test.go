/**
  @author: wangyingjie
  @since: 2023/5/18
  @desc:
**/

package isMatch

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{
			s: "aaa",
			p: "a*aa",
		}, want: true},
		{name: "case1", args: args{
			s: "aaa",
			p: "ab*aa",
		}, want: true},
		{name: "case1", args: args{
			s: "ab",
			p: ".*",
		}, want: true},
		{name: "case1", args: args{
			s: "aa",
			p: "a*",
		}, want: true},
		{name: "case1", args: args{
			s: "ab",
			p: ".*c",
		}, want: false},
		{name: "case1", args: args{
			s: "aab",
			p: "c*a*b",
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}

			if got := isMatchDp(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatchDp() = %v, want %v", got, tt.want)
			}
		})
	}
}
