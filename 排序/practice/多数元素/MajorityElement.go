/**
  @author: wangyingjie
  @since: 2023/2/17
  @desc: https://leetcode.cn/problems/majority-element/
**/

package 多数元素

import "sort"

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[(len(nums)-1)/2]
}
