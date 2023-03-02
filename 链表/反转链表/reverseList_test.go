/**
  @author: wangyingjie
  @since: 2023/3/2
  @desc: //TODO
**/

package 反转链表

import (
	List "algorithm/链表/单链表"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{nums: []int{1, 2, 3, 4, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := List.NewSingleList()
			for _, val := range tt.args.nums {
				list.AddNode(val)
			}
			list.Head = reverseList(list.Head)
			list.Traverse()
		})
	}
}

func Test_reverseListRecursion(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{nums: []int{1, 2, 3, 4, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := List.NewSingleList()
			for _, val := range tt.args.nums {
				list.AddNode(val)
			}
			list.Head = reverseListRecursion(list.Head)
			list.Traverse()
		})
	}
}
