/**
  @author: wangyingjie
  @since: 2023/3/2
  @desc: https://leetcode.cn/problems/longest-continuous-increasing-subsequence/description/
**/

package 最长连续递增序列

import "math"

func findLengthOfLCIS(nums []int) int {
	left, right := 0, 0
	maxLenght := 0
	windowMax := math.MinInt64

	for right < len(nums) {
		if nums[right] > windowMax {
			windowMax = nums[right]
			if maxLenght < right-left+1 {
				maxLenght = right - left + 1
			}
		} else {
			left = right
			windowMax = nums[left]
		}
		right++
	}

	return maxLenght
}
