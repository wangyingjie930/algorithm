/**
  @author: wangyingjie
  @since: 2023/2/21
  @desc
**/

package Rand7实现Rand10

import (
	"fmt"
	"testing"
)

func Test_rand10(t *testing.T) {
	tests := []struct {
		name string
		N    int
	}{
		{name: "case1", N: 1000},
	}
	for _, tt := range tests {
		var test []int
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < tt.N; i++ {
				test = append(test, rand10())
			}
			fmt.Print(test)
		})
	}
}
