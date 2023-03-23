/**
  @author: wangyingjie
  @since: 2023/3/23
  @desc:
**/

package 反转字符串中的单词

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "case1", args: args{s: "the sky is blue"}, want: "blue is sky the"},
		{name: "case1", args: args{s: "  hello world  "}, want: "world hello"},
		{name: "case1", args: args{s: "a good   example"}, want: "example good a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
