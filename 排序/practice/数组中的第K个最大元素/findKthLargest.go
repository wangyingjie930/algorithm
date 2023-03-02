/**
  @author: wangyingjie
  @since: 2023/3/2
  @desc: https://leetcode.cn/problems/kth-largest-element-in-an-array/description/
**/

package 数组中的第K个最大元素

func findKthLargest(nums []int, k int) int {
	if k-1 < 0 || len(nums) < k {
		return -1
	}
	helper(nums, 0, len(nums)-1, k)
	return nums[k-1]
}

func helper(nums []int, start, end int, k int) {
	mid := partition(nums, start, end)
	if mid < k-1 {
		//说明k在mid右侧, 往右区间排序即可排序即可
		helper(nums, mid+1, end, k)
	} else if mid > k-1 {
		helper(nums, start, mid, k)
	}
	return
}

// partition
//  @Description: 实际就是维护i, j这个区间是小于divider的
//  @param nums
//  @param start
//  @param end
//  @return int
func partition(nums []int, start, end int) int {
	divider := nums[start]
	j := start + 1
	for i := start + 1; i <= end; i++ {
		if nums[i] > divider {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	nums[start], nums[j-1] = nums[j-1], nums[start]
	return j - 1
}
