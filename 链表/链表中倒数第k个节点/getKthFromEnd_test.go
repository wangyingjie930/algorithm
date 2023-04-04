/**
  @author: wangyingjie
  @since: 2023/4/3
  @desc:
**/

package 链表中倒数第k个节点

import (
	List "algorithm/链表/单链表"
	"testing"
)

func Test_getKthFromEnd(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{nums: []int{1, 2, 3, 4, 5}, n: 2}},
		{name: "case1", args: args{nums: []int{1, 2, 3, 4, 5, 6}, n: 2}},
		{name: "case1", args: args{nums: []int{1}, n: 1}},
		{name: "case1", args: args{nums: []int{1, 2}, n: 2}},
		{name: "case1", args: args{nums: []int{1, 2, 3}, n: 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := List.NewSingleList()
			for _, val := range tt.args.nums {
				list.AddNode(val)
			}
			list.Head = getKthFromEnd(list.Head, tt.args.n)
			list.Traverse()
		})
	}
}
