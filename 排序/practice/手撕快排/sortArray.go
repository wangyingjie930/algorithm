/**
  @author: wangyingjie
  @since: 2023/3/17
  @desc: https://leetcode.cn/problems/sort-an-array/
**/

package 手撕快排

import "math/rand"

func sortArray(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	helper(nums, 0, len(nums)-1)
	return nums
}

func helper(nums []int, start, end int) {
	if end < 0 || start >= len(nums) || end <= start {
		return
	}
	lt, gt := partition3ways(nums, start, end)
	helper(nums, start, lt)
	helper(nums, gt, end)
}

func partition3ways(nums []int, start, end int) (int, int) {
	index := int(rand.Int31()) % (end - start + 1)
	nums[start], nums[index+start] = nums[index+start], nums[start]

	lt := start    // 表示[start+1, lt]是< nums[start]
	gt := end + 1  // 表示[gt, end]是 > nums[start]
	i := start + 1 // 表示 [lt+1, gt+1]是 = nums[start]
	for i < gt {
		if nums[i] > nums[start] {
			nums[gt-1], nums[i] = nums[i], nums[gt-1]
			gt-- //调换之后, i的位置还未进行判断, 所以不能i++
		} else if nums[i] < nums[start] {
			nums[lt+1], nums[i] = nums[i], nums[lt+1]
			lt++
			//因为[start, lt]维护的是 < nums[start],
			// [lt+1, i-1]维护的是 = nums[start], 所以调换之后原来i的位置的数肯定= nums[start], 所以i可以++
			i++
		} else {
			i++
		}
	}
	nums[lt], nums[start] = nums[start], nums[lt]
	return lt, gt
}
