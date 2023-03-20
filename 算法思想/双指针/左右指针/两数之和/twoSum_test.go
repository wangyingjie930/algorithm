/**
  @author: wangyingjie
  @since: 2023/3/21
  @desc:
**/

package 两数之和

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "case1", args: args{nums: []int{2, 7, 11, 15}, target: 9}, want: []int{0, 1}},
		{name: "case2", args: args{nums: []int{3, 3}, target: 6}, want: []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
