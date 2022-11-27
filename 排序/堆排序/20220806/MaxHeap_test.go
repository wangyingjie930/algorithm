package _0220806

import (
	"math/rand"
	"testing"
)

func TestNewMaxHeap(t *testing.T) {
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
				arr:   append(rand.Perm(1000), rand.Perm(20)...),
			},
		},
		{
			name: "case2",
			args: args{
				arr:  []int{2, 1, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := NewMaxHeap(tt.args.arr)
			/*for heap.count > 0 {
				t.Logf("%+v", heap.ExtractMax())
			}*/
			ret := heap.Sort()
			t.Logf("%+v", ret)
		})
	}
}