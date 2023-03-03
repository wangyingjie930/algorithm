/**
  @author: wangyingjie
  @since: 2023/2/21
  @desc
**/

package 不均匀硬币产生等概率

import (
	"testing"
)

func Test_coin(t *testing.T) {
	tests := []struct {
		name string
		N    int
	}{
		{name: "case1", N: 10000},
	}
	for _, tt := range tests {
		test := []int{0: 0, 1: 0}
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < tt.N; i++ {
				test[averageCoin()]++
			}
			t.Logf("硬币向上的概率: %f", float64(test[0])/float64(tt.N))
			t.Logf("硬币向下的概率: %f", float64(test[1])/float64(tt.N))
		})
	}
}
