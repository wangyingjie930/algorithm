/**
  @author: wangyingjie
  @since: 2023/2/16
  @desc: https://leetcode.cn/problems/first-missing-positive/description/
**/

package 缺失的第一个正数

import "math"

/*func firstMissingPositive(nums []int) int {
	mapper := make(map[int]int, len(nums))
	minNum := math.MaxInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			continue
		}
		mapper[nums[i]] = 1
		if minNum > nums[i] {
			minNum = nums[i]
		}
	}

	if minNum > 1 || minNum == math.MaxInt64 {
		return 1
	}

	i := minNum + 1
	for ; i <= minNum+len(nums); i++ {
		if _, exist := mapper[i]; !exist {
			return i
		}
	}

	return i
}*/

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			nums[i] = len(nums) + 1
		}
	}

	//将[1,2,3...N]所对应的位置标记为负数
	for i := 0; i < len(nums); i++ {
		n := int(math.Abs(float64(nums[i])))
		if n <= len(nums) {
			nums[n-1] = int(math.Abs(float64(nums[n-1]))) * -1
		}
	}

	//相当于遍历了[1,2,3....N]数组
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return i + 1
}
