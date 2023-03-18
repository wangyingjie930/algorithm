/**
  @author: wangyingjie
  @since: 2023/3/18
  @desc:
**/

package 合并两个有序链表

import (
	List "algorithm/链表/单链表"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 []int
		list2 []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{list1: []int{12, 31, 42, 54}, list2: []int{23, 32, 44, 88, 99, 100}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := List.NewSingleList()
			for _, v := range tt.args.list1 {
				list1.AddNode(v)
			}
			list1.Traverse()

			list2 := List.NewSingleList()
			for _, v := range tt.args.list2 {
				list2.AddNode(v)
			}
			list2.Traverse()

			list := List.NewSingleList()
			list.Head = mergeTwoLists(list1.Head, list2.Head)
			list.Traverse()
		})
	}
}

func Test_mergeTwoListsRecursion(t *testing.T) {
	type args struct {
		list1 []int
		list2 []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{list1: []int{12, 31, 42, 54}, list2: []int{23, 32, 44, 88, 99, 100}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := List.NewSingleList()
			for _, v := range tt.args.list1 {
				list1.AddNode(v)
			}
			list1.Traverse()

			list2 := List.NewSingleList()
			for _, v := range tt.args.list2 {
				list2.AddNode(v)
			}
			list2.Traverse()

			list := List.NewSingleList()
			list.Head = mergeTwoListsRecursion(list1.Head, list2.Head)
			list.Traverse()
		})
	}
}
