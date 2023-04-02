/**
  @author: wangyingjie
  @since: 2023/4/3
  @desc:
**/

package 回文链表

import (
	List "algorithm/链表/单链表"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{list: []int{1, 2, 2, 1}}, want: true},
		{name: "case1", args: args{list: []int{1, 2, 3, 2, 1}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := List.NewSingleList()
			for _, v := range tt.args.list {
				list.AddNode(v)
			}
			list.Traverse()

			if got := isPalindrome(list.Head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
