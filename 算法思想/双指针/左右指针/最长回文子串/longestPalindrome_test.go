/**
  @author: wangyingjie
  @since: 2023/2/27
  @desc
**/

package 最长回文子串

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "case1", args: args{s: "babad"}, want: "bab"},
		{name: "case1", args: args{s: "cbbd"}, want: "bb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(longestPalindrome(tt.args.s))
		})
	}
}
