/**
  @author: wangyingjie
  @since: 2023/2/10
  @desc: //TODO
**/

package 在排序数组中查找元素的第一个和最后一个位置

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "case1", args: args{nums: []int{5, 7, 7, 8, 8, 10}, target: 8}, want: []int{3, 4}},
		{name: "case2", args: args{nums: []int{1}, target: 1}, want: []int{0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
