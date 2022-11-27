package _0220806

import (
	"math/rand"
	"testing"
)

func TestNewFixMaxHeap(t *testing.T) {
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
				arr:  []int{2, 1, 300},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := NewFixMaxHeap(tt.args.arr, 1000)
			t.Logf("%+v", *heap)
		})
	}
}
