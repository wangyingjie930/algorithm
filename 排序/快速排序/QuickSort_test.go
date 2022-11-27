package QuickSort

import (
	"math/rand"
	"testing"
)

func TestNormal_Sort(t *testing.T) {
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
				arr:   append(rand.Perm(20), rand.Perm(20)...),
				start: 0,
				end:   39,
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Normal{}
			t.Logf("before: %v", tt.args.arr)
			n.Sort(tt.args.arr, tt.args.start, tt.args.end)
			t.Logf("after: %v", tt.args.arr)
		})
	}
}