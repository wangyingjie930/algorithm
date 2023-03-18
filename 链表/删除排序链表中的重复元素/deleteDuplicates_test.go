/**
  @author: wangyingjie
  @since: 2023/3/18
  @desc:
**/

package 删除排序链表中的重复元素

import (
	List "algorithm/链表/单链表"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{head: []int{1, 1, 2, 2, 3, 3, 4, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := List.NewSingleList()
			for _, v := range tt.args.head {
				list1.AddNode(v)
			}
			list1.Traverse()

			list1.Head = deleteDuplicates(list1.Head)
			list1.Traverse()
		})
	}
}
