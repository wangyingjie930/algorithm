/**
  @author: wangyingjie
  @since: 2023/2/15
  @desc: //TODO
**/

package 一手顺子

import "testing"

func Test_isNStraightHand(t *testing.T) {
	type args struct {
		hand      []int
		groupSize int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{hand: []int{1, 2, 3, 6, 2, 3, 4, 7, 8}, groupSize: 3}, want: true},
		{name: "case1", args: args{hand: []int{3, 4, 7, 4, 6, 3, 6, 5, 2, 8}, groupSize: 2}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNStraightHand(tt.args.hand, tt.args.groupSize); got != tt.want {
				t.Errorf("isNStraightHand() = %v, want %v", got, tt.want)
			}
		})
	}
}
