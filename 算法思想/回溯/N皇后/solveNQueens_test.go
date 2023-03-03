/**
  @author: wangyingjie
  @since: 2023/2/25
  @desc
**/

package N皇后

import (
	"fmt"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case", args: args{n: 4}},
		{name: "case", args: args{n: 3}},
		{name: "case", args: args{n: 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solveNQueens(tt.args.n)
			fmt.Println(got)
		})
	}
}
