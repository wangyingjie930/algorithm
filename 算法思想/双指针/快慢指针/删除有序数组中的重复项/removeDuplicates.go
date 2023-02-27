/**
  @author: wangyingjie
  @since: 2023/2/27
  @desc: https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
**/

package 删除有序数组中的重复项

func removeDuplicates(nums []int) int {
	slow := 0
	for quick := slow + 1; quick < len(nums); quick++ {
		if nums[quick] != nums[slow] {
			slow++
			nums[slow] = nums[quick]
		}
	}
	return slow + 1
}
