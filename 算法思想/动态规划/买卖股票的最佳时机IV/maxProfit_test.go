/**
  @author: wangyingjie
  @since: 2023/2/28
  @desc
**/

package 买卖股票的最佳时机IV

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		k      int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{k: 2, prices: []int{2, 4, 1}}, want: 2},
		{name: "case1", args: args{k: 2, prices: []int{3, 2, 6, 5, 0, 3}}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
