/**
  @author: wangyingjie
  @since: 2023/2/9
  @desc: https://leetcode.cn/problems/search-in-rotated-sorted-array-ii/
**/

package 搜索旋转排序数组2

// search
//  @Description:  参考https://imageslr.com/2020/03/06/leetcode-33.html
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
			// 这个条件说明从mid到right这段路是单调递增的
			if nums[mid] < target && target <= nums[right] {
				//因为是从mid到right单调递增, 所以假设将target在这个区间中
				//接下来使用二分法的那个模板, 那么left必须向右移动来维护左区间为 x < target的左区间
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			// 这个条件说明从left到mid这段路是单调递增的
			if nums[mid] >= target && target >= nums[left] {
				//接下来使用二分法的那个模板, 那么right必须向左移动来维护右区间为 x >= target的右区间
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return false
}
