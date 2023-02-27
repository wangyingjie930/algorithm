/**
  @author: wangyingjie
  @since: 2023/2/23
  @desc: //TODO
**/

package 爬楼梯

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{n: 2}, want: 2},
		{name: "case1", args: args{n: 3}, want: 3},
		{name: "case1", args: args{n: 5}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairsDp(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
