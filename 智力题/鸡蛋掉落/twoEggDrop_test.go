/**
  @author: wangyingjie
  @since: 2023/4/2
  @desc:
**/

package 鸡蛋掉落

import "testing"

func Test_twoEggDrop(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{n: 2}, want: 2},
		{name: "case1", args: args{n: 100}, want: 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoEggDrop(tt.args.n); got != tt.want {
				t.Errorf("twoEggDrop() = %v, want %v", got, tt.want)
			}
		})
	}
}
