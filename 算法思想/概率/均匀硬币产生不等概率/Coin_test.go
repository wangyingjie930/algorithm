/**
  @author: wangyingjie
  @since: 2023/2/21
  @desc
**/

package 均匀硬币产生不等概率

import "testing"

func Test_newCoin(t *testing.T) {
	tests := []struct {
		name string
		N    int
	}{
		{name: "case1", N: 10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := []int{0: 0, 1: 0}
			for i := 0; i < tt.N; i++ {
				test[newCoin1divide4()]++
			}
			t.Logf("硬币向上的概率: %f", float64(test[1])/float64(tt.N))
			t.Logf("硬币向下的概率: %f", float64(test[0])/float64(tt.N))
		})
	}
}

func Test_newCoin1divide3(t *testing.T) {
	tests := []struct {
		name string
		N    int
	}{
		{name: "case1", N: 10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := []int{0: 0, 1: 0}
			for i := 0; i < tt.N; i++ {
				test[newCoin1divide3()]++
			}
			t.Logf("硬币向上的概率: %f", float64(test[1])/float64(tt.N))
			t.Logf("硬币向下的概率: %f", float64(test[0])/float64(tt.N))
		})
	}
}

func Test_newCoin3divide10(t *testing.T) {
	tests := []struct {
		name string
		N    int
	}{
		{name: "case1", N: 10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := []int{0: 0, 1: 0}
			for i := 0; i < tt.N; i++ {
				test[newCoin3divide10()]++
			}
			t.Logf("硬币向上的概率: %f", float64(test[1])/float64(tt.N))
			t.Logf("硬币向下的概率: %f", float64(test[0])/float64(tt.N))
		})
	}
}
