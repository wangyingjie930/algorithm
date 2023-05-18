/**
  @author: wangyingjie
  @since: 2023/5/18
  @desc:
**/

package findKthNumber

import "testing"

func Test_findKthNumber(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{
			n: 13,
			k: 2,
		}, want: 10},
		{name: "case1", args: args{
			n: 1,
			k: 1,
		}, want: 1},
		{name: "case1", args: args{
			n: 2,
			k: 2,
		}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthNumber(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("findKthNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
