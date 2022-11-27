package 冒泡排序

import (
	"math/rand"
	"testing"
)

func TestSort(t *testing.T) {
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
				arr:   append(rand.Perm(2000), rand.Perm(20)...),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort(tt.args.arr)
			t.Logf("%v", tt.args.arr)
		})
	}
}