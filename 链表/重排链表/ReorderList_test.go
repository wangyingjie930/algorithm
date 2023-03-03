/**
  @author: wangyingjie
  @since: 2023/2/10
  @desc
**/

package 重排链表

import "testing"
import List "algorithm/链表/单链表"

func Test_reorderList(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{nums: []int{1, 2, 3, 4, 5}}},
		{name: "case1", args: args{nums: []int{1, 2, 3, 4, 5, 6}}},
		{name: "case1", args: args{nums: []int{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := List.NewSingleList()
			for _, val := range tt.args.nums {
				list.AddNode(val)
			}
			reorderList(list.Head)
			list.Traverse()
		})
	}
}
