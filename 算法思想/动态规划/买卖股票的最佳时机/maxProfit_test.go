/**
  @author: wangyingjie
  @since: 2023/3/2
  @desc
**/

package 买卖股票的最佳时机

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{prices: []int{7, 1, 5, 3, 6, 4}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
