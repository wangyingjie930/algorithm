/**
  @author: wangyingjie
  @since: 2023/2/10
  @desc: https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
**/

package 在排序数组中查找元素的第一个和最后一个位置

const LEFT_RANGE, RIGHT_RANGE = 1, 2

func searchRange(nums []int, target int) []int {
	ret := []int{-1, -1}
	if len(nums) == 0 {
		return ret
	}

	start := binarySearch(nums, target, LEFT_RANGE)
	end := binarySearch(nums, target, RIGHT_RANGE) - 1

	if start <= end && end < len(nums) && nums[start] == target && nums[end] == target {
		return []int{start, end}
	}
	return ret
}

// binarySearch
//  @Description: 使用二分查找法
//  @param nums
//  @param target
//  @param rangeType 根据类型获取查找最左/右的target
//  @return int
func binarySearch(nums []int, target int, rangeType int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if rangeType == LEFT_RANGE {
			//这个范围定义右区间的范围可以覆盖到最左侧的target
			if nums[mid] >= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if rangeType == RIGHT_RANGE {
			//这个范围定义右区间的范围可以覆盖最右侧的target再往后的一个数
			if nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return left
}
