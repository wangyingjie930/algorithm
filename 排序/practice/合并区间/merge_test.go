/**
  @author: wangyingjie
  @since: 2023/3/19
  @desc:
**/

package 合并区间

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "case1", args: args{intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}}, want: [][]int{{1, 6}, {8, 10}, {15, 18}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
