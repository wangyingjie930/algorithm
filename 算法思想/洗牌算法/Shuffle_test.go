/**
  @author: wangyingjie
  @since: 2023/2/14
  @desc
**/

package 洗牌算法

import (
	"fmt"
	"testing"
)

func Test_shuffle(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case2", args: args{nums: []int{1, 2, 31, 411, 5}}},
		{name: "case2", args: args{nums: []int{11, 21, 3, 4, 5}}},
		{name: "case2", args: args{nums: []int{1, 2, 3, 411, 5}}},
		{name: "case2", args: args{nums: []int{11, 12, 31, 4, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shuffle(tt.args.nums)
			fmt.Println(tt.args.nums)
		})
	}
}
