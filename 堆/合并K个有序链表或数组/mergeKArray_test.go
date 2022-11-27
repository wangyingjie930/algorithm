package 合并K个有序链表或数组

import (
	"testing"
)

func Test_mergeKArray(t *testing.T) {
	ret := mergeKArray([][]int{
		{9, 6, 4, 2, 1},
		{19, 8, 5, 3, 1},
	})
	t.Logf("%+v", ret)
}
