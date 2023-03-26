/**
  @author: wangyingjie
  @since: 2023/3/27
  @desc:
**/

package 最长重复子数组

import "testing"

func Test_findLength(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{nums1: []int{1, 2, 3, 2, 1}, nums2: []int{3, 2, 1, 4, 7}}, want: 3},
		{name: "case1", args: args{nums1: []int{0, 0, 0, 0, 0}, nums2: []int{0, 0, 0, 0, 0}}, want: 5},
		{name: "case1", args: args{nums1: []int{0, 1, 1, 1, 1}, nums2: []int{1, 0, 1, 0, 1}}, want: 2},
		{name: "case1", args: args{nums1: []int{5, 14, 53, 80, 48}, nums2: []int{50, 47, 3, 80, 83}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLength(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
