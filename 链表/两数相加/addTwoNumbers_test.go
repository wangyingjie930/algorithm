/**
  @author: wangyingjie
  @since: 2023/3/22
  @desc:
**/

package 两数相加

import (
	List "algorithm/链表/单链表"
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "case1", args: args{l1: []int{2, 4, 3}, l2: []int{5, 6, 4}}, want: []int{7, 0, 8}},
		{name: "case1", args: args{l1: []int{0}, l2: []int{0}}, want: []int{0}},
		{name: "case1", args: args{l1: []int{9, 9, 9, 9, 9, 9, 9}, l2: []int{9, 9, 9, 9}}, want: []int{8, 9, 9, 9, 0, 0, 0, 1}},
		{name: "case1", args: args{l1: []int{2, 4, 9}, l2: []int{5, 6, 4, 9}}, want: []int{7, 0, 4, 0, 1}},
		{name: "case1", args: args{l1: []int{5}, l2: []int{5}}, want: []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := List.NewSingleList()
			for _, v := range tt.args.l1 {
				l1.AddNode(v)
			}
			l1.Traverse()

			l2 := List.NewSingleList()
			for _, v := range tt.args.l2 {
				l2.AddNode(v)
			}
			l2.Traverse()

			got := addTwoNumbers(l1.Head, l2.Head)
			var res []int
			for got != nil {
				res = append(res, got.Val)
				got = got.Next
			}
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", res, tt.want)
			}
		})
	}
}
