/**
  @author: wangyingjie
  @since: 2023/3/18
  @desc:
**/

package 环形链表II

import (
	List "algorithm/链表/单链表"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head       []int
		beginValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{head: []int{1, 2, 3, 4, 5, 6, 7}, beginValue: 3}, want: 3},
		{name: "case2", args: args{
			head: []int{9863, -7835, -2576, -1390, 1154, 13654, -6793, 14489, 14739, 15465, 12843, 14656, 9998, -4416, 10104, 4804, 1556, 9350, 5791, 8999, -4762, 12242, 13264, -7520, 10069, 4058, -2512}, beginValue: -1}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := List.NewSingleList()
			var beginNode *List.Node
			for _, v := range tt.args.head {
				list.AddNode(v)
				if v == tt.args.beginValue && tt.args.beginValue != -1 {
					beginNode = list.Tail
				}
			}
			if beginNode != nil {
				list.Tail.Next = beginNode
			}
			list.Traverse()

			t.Log(detectCycle(list.Head))
		})
	}
}
