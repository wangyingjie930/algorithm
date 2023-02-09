/**
  @author: wangyingjie
  @since: 2023/2/9
  @desc: https://leetcode.cn/problems/search-in-rotated-sorted-array-ii/
**/

package 搜索旋转排序数组2

// search
//  @Description: 说明: 目前的解法边界条件有问题还需梳理
//  @param nums
//  @param target
//  @return bool
func search(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return true
		}

		if nums[left] == nums[mid] && nums[right] == nums[mid] {
			left++
			right--
		} else if nums[mid] <= nums[right] {
			if target >= nums[right] || target < nums[mid] {
				right = mid - 1
			} else {
				left = mid
			}
		} else {
			if target <= nums[right] || target > nums[mid] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}
	return false
}
