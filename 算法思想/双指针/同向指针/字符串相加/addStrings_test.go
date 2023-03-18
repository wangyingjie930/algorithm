/**
  @author: wangyingjie
  @since: 2023/3/18
  @desc:
**/

package 字符串相加

import "testing"

func Test_addStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "case1", args: args{num1: "123", num2: "456"}, want: "579"},
		{name: "case2", args: args{num1: "456", num2: "77"}, want: "533"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
