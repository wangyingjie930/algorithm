/**
  @author: wangyingjie
  @since: 2023/4/2
  @desc:
**/

package 两两交换链表中的节点

import (
	List "algorithm/链表/单链表"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{head: []int{1, 2, 3, 4}}},
		{name: "case1", args: args{head: []int{}}},
		{name: "case1", args: args{head: []int{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := List.NewSingleList()
			for _, v := range tt.args.head {
				list1.AddNode(v)
			}
			list1.Traverse()

			list1.Head = swapPairs(list1.Head)
			list1.Traverse()
		})
	}
}

func Test_swapPairsIteration(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{head: []int{1, 2, 3, 4}}},
		{name: "case1", args: args{head: []int{}}},
		{name: "case1", args: args{head: []int{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := List.NewSingleList()
			for _, v := range tt.args.head {
				list1.AddNode(v)
			}
			list1.Traverse()

			list1.Head = swapPairsIteration(list1.Head)
			list1.Traverse()
		})
	}
}
