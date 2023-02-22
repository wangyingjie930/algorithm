/**
  @author: wangyingjie
  @since: 2023/2/8
  @desc: 寻找旋转排序数组中的最小值
  @see: https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/description/
**/

package 寻找旋转排序数组中的最小值

/*func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	start := 0
	end := len(nums)
	return int(math.Min(float64(findMin(nums[start:(start+end)/2])), float64(findMin(nums[(start+end)/2:]))))
}*/

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	//二分: 这里以闭区间作为边界限制
	for left <= right {
		pivot := (left + right) / 2
		// 如果中间比右边界更小, 证明右边界集中了更多的较高值, 需向左边界缩小范围去查找最小值
		//不可以<=, 当找到最小值后nums[pivot]肯定等于nums[right], 会进入死循环
		if nums[pivot] < nums[right] {
			// 不可以pivot-1, 当[3,1,2]这种情况直接把1排除出去了
			right = pivot
		} else {
			left = pivot + 1
		}
	}
	//二分: 因为当left==right时, nums[pivot]肯定等于nums[right], left肯定变了, 取right就行
	return nums[right]
}
