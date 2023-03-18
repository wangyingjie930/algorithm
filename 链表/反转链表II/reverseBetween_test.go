/**
  @author: wangyingjie
  @since: 2023/3/18
  @desc:
**/

package 反转链表II

import (
	List "algorithm/链表/单链表"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{head: []int{1, 2, 3, 4, 5}, left: 2, right: 4}},
		{name: "case2", args: args{head: []int{5}, left: 1, right: 1}},
		{name: "case3", args: args{head: []int{3, 5}, left: 1, right: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := List.NewSingleList()
			for _, v := range tt.args.head {
				list.AddNode(v)
			}
			list.Head = reverseBetween(list.Head, tt.args.left, tt.args.right)
			list.Traverse()
		})
	}
}
