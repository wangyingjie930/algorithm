/**
  @author: wangyingjie
  @since: 2023/4/3
  @desc:
**/

package 比较版本号

import "testing"

func Test_compareVersion(t *testing.T) {
	type args struct {
		version1 string
		version2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{version1: "1.01", version2: "1.001"}, want: 0},
		{name: "case1", args: args{version1: "0.1", version2: "1.1"}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareVersion(tt.args.version1, tt.args.version2); got != tt.want {
				t.Errorf("compareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
