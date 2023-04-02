/**
  @author: wangyingjie
  @since: 2023/4/2
  @desc:
**/

package 单词拆分

import "testing"

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{s: "leetcode", wordDict: []string{"leet", "code"}}, want: true},
		{name: "case1", args: args{s: "applepenapple", wordDict: []string{"apple", "pen"}}, want: true},
		{name: "case1", args: args{s: "catsandog", wordDict: []string{"cats", "dog", "sand", "and", "cat"}}, want: false},
		{name: "case1", args: args{s: "a", wordDict: []string{"a"}}, want: true},
		{name: "case1", args: args{s: "ab", wordDict: []string{"a", "b"}}, want: true},
		{name: "case1", args: args{s: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", wordDict: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wordBreakDp(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{s: "leetcode", wordDict: []string{"leet", "code"}}, want: true},
		{name: "case1", args: args{s: "applepenapple", wordDict: []string{"apple", "pen"}}, want: true},
		{name: "case1", args: args{s: "catsandog", wordDict: []string{"cats", "dog", "sand", "and", "cat"}}, want: false},
		{name: "case1", args: args{s: "a", wordDict: []string{"a"}}, want: true},
		{name: "case1", args: args{s: "ab", wordDict: []string{"a", "b"}}, want: true},
		{name: "case1", args: args{s: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", wordDict: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreakDp(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreakDp() = %v, want %v", got, tt.want)
			}
		})
	}
}
