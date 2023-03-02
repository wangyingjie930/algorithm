/**
  @author: wangyingjie
  @since: 2023/2/18
  @desc: https://leetcode.cn/problems/search-in-rotated-sorted-array/description/
  @see https://imageslr.com/2020/03/15/binary-search.html

**/

package 搜索旋转排序数组

func search(nums []int, target int) int {
	index := helper(nums, 0, len(nums)-1, target)
	if index < 0 || len(nums) <= index || nums[index] != target {
		return -1
	}
	return index
}

func helper(nums []int, left, right int, target int) int {
	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			//这里是因为有报错加的, 由于二分的区间有变化所以对这个情况特殊处理
			return mid
		}

		if nums[mid] <= nums[right] {
			//从mid到right单调递增
			if target > nums[mid] && target <= nums[right] {
				//由于mid到right是单调递增, 那么如果target处于 nums[mid] < target <= num[right]就可以使用二分的模板了(符合有序的前提)
				//target > nums[mid]刚好是符合二分左区间的条件左区间扩大
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			//从left到mid单调递增
			if target <= nums[mid] && nums[left] <= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return left
}
