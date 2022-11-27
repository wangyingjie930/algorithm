package 计数排序

import (
	"math/rand"
	"testing"
)

func TestCountingSort(t *testing.T) {
	type args struct {
		arr   []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				arr:   append(rand.Perm(20), []int{2,3,4,4,4,4,4,33,2,2,6,7,2}...),
				start: 0,
				end:   19,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("before: %v", tt.args.arr)
			CountingSort(tt.args.arr)
			t.Logf("after: %v", tt.args.arr)
		})
	}
}
