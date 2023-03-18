/**
  @author: wangyingjie
  @since: 2023/3/19
  @desc: https://leetcode.cn/problems/merge-intervals/solutions/
**/

package 合并区间

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	//先通过排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	//后比对结束区间是否交集
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		lastIntervals := res[len(res)-1]
		if lastIntervals[1] >= intervals[i][0] {
			lastIntervals[1] = int(math.Max(float64(intervals[i][1]), float64(lastIntervals[1])))
		} else {
			res = append(res, intervals[i])
		}
	}

	return res
}
