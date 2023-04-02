/**
  @author: wangyingjie
  @since: 2023/4/2
  @desc:
**/

package 最长公共前缀

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "case1", args: args{strs: []string{"flower", "flow", "flight"}}, want: "fl"},
		{name: "case1", args: args{strs: []string{"dog", "racecar", "car"}}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
