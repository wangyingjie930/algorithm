/**
  @author: wangyingjie
  @since: 2023/2/8
  @desc
**/

package K个一组翻转链表

import (
	List "algorithm/链表/单链表"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{nums: []int{1, 2, 3, 4, 5}, k: 2}},
		{name: "case1", args: args{nums: []int{1, 2, 3, 4, 5}, k: 3}},
		{name: "case1", args: args{nums: []int{}, k: 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := List.NewSingleList()
			for _, v := range tt.args.nums {
				list1.AddNode(v)
			}
			list1.Traverse()

			list1.Head = reverseKGroup(list1.Head, tt.args.k)
			list1.Traverse()
		})
	}
}

func Test_reverseKGroupIteration(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{nums: []int{1, 2, 3, 4, 5}, k: 2}},
		{name: "case1", args: args{nums: []int{1, 2, 3, 4, 5}, k: 3}},
		{name: "case1", args: args{nums: []int{}, k: 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := List.NewSingleList()
			for _, v := range tt.args.nums {
				list1.AddNode(v)
			}
			list1.Traverse()

			list1.Head = reverseKGroupIteration(list1.Head, tt.args.k)
			list1.Traverse()
		})
	}
}
