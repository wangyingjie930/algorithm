/**
  @author: wangyingjie
  @since: 2023/9/17
  @desc:
**/

package InsertSort

import (
	"math/rand"
	"testing"
)

func TestInsertSort1(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				arr: append(rand.Perm(2000), rand.Perm(20)...),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertSort(tt.args.arr)
			t.Logf("%v", tt.args.arr)
		})
	}
}
