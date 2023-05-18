/**
  @author: wangyingjie
  @since: 2023/5/18
  @desc:
**/

package calculate

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{s: "1+2*3+4"}, want: 11},
		{name: "case1", args: args{s: "(1+(4+5+2)-3)+(6+8)"}, want: 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
