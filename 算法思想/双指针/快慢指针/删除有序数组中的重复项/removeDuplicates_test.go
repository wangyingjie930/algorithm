/**
  @author: wangyingjie
  @since: 2023/2/27
  @desc
**/

package 删除有序数组中的重复项

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{[]int{1, 1, 2}}, want: 2},
		{name: "case1", args: args{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
			t.Log(tt.args.nums)
		})
	}
}
