/**
  @author: wangyingjie
  @since: 2023/3/11
  @desc:
**/

package MapReduce

import (
	"fmt"
	"testing"
)

func Test_runMapReduce(t *testing.T) {
	type args struct {
		input   []string
		mapper  func(string) []KeyValue
		reducer func(string, []string) string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{input: []string{
			"hello world",
			"world world",
			"world hello",
			"go go go",
			"go world",
			"hello world",
			"world world",
			"world hello",
			"go go go",
			"go world",
			"hello world",
			"world world",
			"world hello asda adaw12 fasd 12",
			"go go go",
			"go world",
		}, mapper: mapper, reducer: reducer}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := runMapReduce(tt.args.input, tt.args.mapper, tt.args.reducer)
			fmt.Println(got)
		})
	}
}
