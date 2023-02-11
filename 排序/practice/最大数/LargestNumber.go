/**
  @author: wangyingjie
  @since: 2023/2/11
  @desc: https://leetcode.cn/problems/largest-number/
**/

package 最大数

import (
	"sort"
	"strconv"
)

/**
输入：nums = [3,30,34,5,9]
输出："9534330"
*/
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		sa, sb := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		//将两数拼接, 长度相等
		//比较拼接后的大小, 大的字符串则返回开头的数
		sa = sa + sb
		sb = sb + sa
		return sa > sb
	})
	if nums[0] == 0 {
		//对于[0, 0, 0]的处理
		return "0"
	}
	ret := ""
	for _, num := range nums {
		ret = ret + strconv.Itoa(num)
	}
	return ret
}
