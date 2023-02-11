/**
  @author: wangyingjie
  @since: 2023/2/10
  @desc: https://leetcode.cn/problems/sort-colors/
**/

package 颜色分类

func sortColors(nums []int) {
	sortColor(nums, 0, len(nums)-1)
}

// sortColor
//  @Description: 因为题目规定原地排序, 所以采用快速排序
//  @param nums
//  @param start
//  @param end
func sortColor(nums []int, start, end int) {
	if start >= end {
		return
	}

	mid := partition(nums, start, end)
	sortColor(nums, start, mid-1)
	sortColor(nums, mid+1, end)
}

func partition(nums []int, start, end int) int {
	judge := nums[start]
	i := start
	for j := start + 1; j <= end; j++ {
		if nums[j] <= judge {
			nums[i+1], nums[j] = nums[j], nums[i+1]
			i++
		}
	}
	nums[i], nums[start] = nums[start], nums[i]
	return i
}
