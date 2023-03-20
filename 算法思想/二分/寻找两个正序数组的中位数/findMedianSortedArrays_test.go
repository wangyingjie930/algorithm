/**
  @author: wangyingjie
  @since: 2023/3/21
  @desc:
**/

package 寻找两个正序数组的中位数

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "case1", args: args{nums1: []int{1, 3}, nums2: []int{2}}, want: 2},
		{name: "case1", args: args{nums1: []int{1, 2}, nums2: []int{3, 4}}, want: 2.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
