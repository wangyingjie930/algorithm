/**
  @author: wangyingjie
  @since: 2023/2/8
  @desc: 寻找旋转排序数组中的最小值
  @see: https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/description/
**/

package 寻找旋转排序数组中的最小值

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < nums[right] {
			//表明在这个区间单调递增
			right = mid //不-1是因为可能mid是最小值
		} else {
			left = mid + 1
		}
	}
	return nums[right]
}
