/**
  @author: wangyingjie
  @since: 2023/3/26
  @desc:
**/

package 组合总和

import (
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "case1", args: args{candidates: []int{2, 3, 6, 7}, target: 7}, want: [][]int{{2, 2, 3}, {7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
