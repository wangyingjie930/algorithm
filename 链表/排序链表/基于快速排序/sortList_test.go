/**
  @author: wangyingjie
  @since: 2023/3/22
  @desc:
**/

package 基于归并排序

import (
	List "algorithm/链表/单链表"
	"testing"
)

func Test_sortList(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{list: []int{3, 2, 1, 4, 5, 1}}},
		{name: "case1", args: args{list: []int{3}}},
		{name: "case1", args: args{list: []int{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := List.NewSingleList()
			for _, v := range tt.args.list {
				list.AddNode(v)
			}
			list.Traverse()

			list.Head = sortList(list.Head)
			list.Traverse()
		})
	}
}
