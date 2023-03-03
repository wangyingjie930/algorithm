/**
  @author: wangyingjie
  @since: 2023/2/17
  @desc
**/

package 相交链表

import (
	List "algorithm/链表/单链表"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		numA         []int
		numB         []int
		intersectVal int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{numA: []int{4, 1, 8, 4, 5}, numB: []int{5, 6, 1, 8, 4, 5}, intersectVal: 8}, want: 8},
		{name: "case1", args: args{numA: []int{4, 1, 8, 4, 5}, numB: []int{5, 6, 1, 8, 4, 5}, intersectVal: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listA := List.NewSingleList()
			listB := List.NewSingleList()
			var intersetNode *List.Node
			for i := 0; i < len(tt.args.numA); i++ {
				listA.AddNode(tt.args.numA[i])
				if tt.args.numA[i] == tt.args.intersectVal {
					intersetNode = listA.Tail
				}
			}
			for i := 0; i < len(tt.args.numB); i++ {
				if tt.args.numB[i] != tt.args.intersectVal {
					listB.AddNode(tt.args.numB[i])
				} else {
					listB.Tail.Next = intersetNode
					listB.Tail = listA.Tail
					break
				}
			}

			if got := getIntersectionNode(listA.Head, listB.Head); got.Val != tt.want {
				t.Errorf("getIntersectionNode() = %v, want %v", got.Val, tt.want)
			}
		})
	}
}
