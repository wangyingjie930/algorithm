package _0220806

import (
	"container/heap"
	"testing"
)

func TestNewPriorityQueue(t *testing.T) {
	type args struct {
		items map[string]int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{
			map[string]int{"aa": 1, "bb": 100, "cc": 23, "dd": 50, "ee": 25},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := NewPriorityQueue(tt.args.items)
			pq.Println()

			for pq.Len() > 0 {
				t.Logf("%+v", heap.Pop(pq))
			}
		})
	}
}

func Test_mergeKSortList(t *testing.T) {
	type args struct {
		lists []map[string]int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{
			[]map[string]int{
				{"aa": 1, "bb": 10, "cc": 223, "dd": 530, "ee": 215},
				{"a1": 1, "b1": 100, "c1": 23, "d1": 50, "e1": 25},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := mergeKSortList(tt.args.lists)
			t.Logf("%+v", pq)
		})
	}
}