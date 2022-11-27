package QuickSort

import (
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {
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
				arr:   rand.Perm(20),
				start: 0,
				end:   19,
			},
		},
		{
			name: "case2",
			args: args{
				arr:   rand.Perm(20),
				start: 0,
				end:   19,
			},
		},
		{
			name: "case3",
			args: args{
				arr:   rand.Perm(200),
				start: 0,
				end:   199,
			},
		},
		{
			name: "case3",
			args: args{
				arr:   []int{49, 34},
				start: 0,
				end:   199,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("before: %v", tt.args.arr)
			QuickSort(tt.args.arr)
			t.Logf("after: %v", tt.args.arr)
		})
	}
}
