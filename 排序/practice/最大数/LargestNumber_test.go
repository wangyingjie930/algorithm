/**
  @author: wangyingjie
  @since: 2023/2/11
  @desc
**/

package 最大数

import "testing"

func Test_largestNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "case1", args: args{nums: []int{3, 30, 34, 5, 9}}, want: "9534330"},
		{name: "case1", args: args{nums: []int{10, 2}}, want: "210"},
		{name: "case1", args: args{nums: []int{111311, 1113}}, want: "1113111311"},
		{name: "case1", args: args{nums: []int{0, 0, 0, 0}}, want: "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.args.nums); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
