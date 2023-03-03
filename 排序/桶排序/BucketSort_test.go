/**
  @author: wangyingjie
  @since: 2022/11/27
  @desc
**/

package 桶排序

import "testing"

func Test_bucketSort(t *testing.T) {
	type args struct {
		arr   []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				arr:   []int{111111, 22323, 4534, 2322, 4365, 25254, 512, 24321, 532, 7457, 7443, 2342},
				start: 0,
				end:   19,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("before: %v", tt.args.arr)
			bucketSort(tt.args.arr, 10)
			t.Logf("after: %v", tt.args.arr)
		})
	}
}
