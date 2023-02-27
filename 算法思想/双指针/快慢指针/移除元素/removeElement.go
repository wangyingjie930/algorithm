/**
  @author: wangyingjie
  @since: 2023/2/27
  @desc: https://leetcode.cn/problems/remove-element/
**/

package 移除元素

func removeElement(nums []int, val int) int {
	slow := 0
	for quick := 0; quick < len(nums); quick++ {
		if nums[quick] != val {
			nums[slow] = nums[quick]
			slow++
		}
	}
	return slow
}
