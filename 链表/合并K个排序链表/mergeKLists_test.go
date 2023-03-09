/**
  @author: wangyingjie
  @since: 2023/3/10
  @desc:
**/

package 合并K个排序链表

import (
	List "algorithm/链表/单链表"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{lists: [][]int{{2, 33, 55, 66}, {12, 31, 42, 54}, {23, 32, 44, 88, 99, 100}}}},
		{name: "case2", args: args{lists: [][]int{{-8, 2, 4}, {-9, -9, -9, -9, -8, -5, -3, -2, 1}, {-9, -7, -5, -3, -3, -1, 0, 3, 4}, {-9, -7, -6, -4, -2, -1, 3, 4}, {-10, -3, 0}, {-9, 0, 4}, {-8, -8}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			lists := make([]*List.Node, len(tt.args.lists))
			for i := 0; i < len(tt.args.lists); i++ {
				list := List.NewSingleList()
				for j := 0; j < len(tt.args.lists[i]); j++ {
					list.AddNode(tt.args.lists[i][j])
				}
				lists[i] = list.Head
				list.Traverse()
			}

			node := mergeKLists(lists)
			list := &List.List{Head: node}
			list.Traverse()
		})
	}
}
