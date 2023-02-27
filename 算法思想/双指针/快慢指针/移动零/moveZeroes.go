/**
  @author: wangyingjie
  @since: 2023/2/27
  @desc: https://leetcode.cn/problems/move-zeroes/
**/

package 移动零

func moveZeroes(nums []int) {
	slow := 0
	for quick := 0; quick < len(nums); quick++ {
		if nums[quick] != 0 {
			nums[slow] = nums[quick]
			slow++
		}
	}
	for ; slow < len(nums); slow++ {
		nums[slow] = 0
	}
}
