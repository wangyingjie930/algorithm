/**
  @author: wangyingjie
  @since: 2023/3/6
  @desc:
**/

package 分布式锁

import "testing"

func Test_newLock(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{i: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newLock(tt.args.i)
		})
	}
}
