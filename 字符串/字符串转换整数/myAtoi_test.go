/**
  @author: wangyingjie
  @since: 2023/3/23
  @desc:
**/

package 字符串转换整数

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{s: "1212121"}, want: 1212121},
		{name: "case1", args: args{s: "        -1212121"}, want: -1212121},
		{name: "case1", args: args{s: "        -1212 hello 121"}, want: -1212},
		{name: "case1", args: args{s: "words and 987"}, want: 0},
		{name: "case1", args: args{s: "-91283472332"}, want: -2147483648},
		{name: "case1", args: args{s: "91283472332"}, want: 2147483647},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
