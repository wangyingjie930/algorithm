/**
  @author: wangyingjie
  @since: 2023/2/27
  @desc: //TODO
**/

package 移除元素

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{nums: []int{3, 2, 2, 3}, val: 3}, want: 2},
		{name: "case2", args: args{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
			t.Log(tt.args.nums)
		})
	}
}
