/**
  @author: wangyingjie
  @since: 2023/3/28
  @desc: https://leetcode.cn/problems/find-peak-element/description/
**/

package 寻找峰值

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if (mid-1 < 0 || nums[mid] > nums[mid-1]) && (mid+1 > len(nums)-1 || nums[mid] > nums[mid+1]) {
			return mid
		}
		if mid+1 < len(nums) && nums[mid] < nums[mid+1] {
			//只要右边的第一个数比nums[mid]大, 那么一定可以在右边找到峰值
			// 要么真的有峰值
			// 要么就是单调递增, 由于题意说的是nums[n] = -∞, 那么数组的最后一个就是峰值
			left = mid + 1
		} else if mid-1 >= 0 && nums[mid] < nums[mid-1] {
			right = right - 1
		} else {
			left++
			right--
		}
	}
	return left
}
