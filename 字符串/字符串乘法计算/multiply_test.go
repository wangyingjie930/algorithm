/**
  @author: wangyingjie
  @since: 2023/3/24
  @desc:
**/

package 字符串乘法计算

import "testing"

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "case1", args: args{num1: "123", num2: "45"}, want: "5535"},
		{name: "case1", args: args{num1: "4444", num2: "477"}, want: "2119788"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
