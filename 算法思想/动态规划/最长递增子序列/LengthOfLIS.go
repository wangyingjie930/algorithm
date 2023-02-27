/**
  @author: wangyingjie
  @since: 2023/2/9
  @desc: https://leetcode.cn/problems/longest-increasing-subsequence/
**/

package 最长递增子序列

import "math"

func lengthOfLIS(nums []int) int {
	maxSum := 0
	for i, _ := range nums {
		//遍历以i为结束节点的排序序列最大长度
		maxSum = int(math.Max(float64(itemMaxSum(nums, i)), float64(maxSum)))
	}
	return maxSum
}

// itemMaxSum
//  @Description: 计算以i为结束节点的排序序列最大长度
//  @param nums
//  @param i
//  @return int
func itemMaxSum(nums []int, i int) int {
	if i == 0 {
		return 1
	}

	sum := 1
	for j := i - 1; j >= 0; j-- {
		if nums[j] >= nums[i] {
			continue
		}
		//获取nums[i]之前的所有元素
		//依次进行比对, 选取最长的并加上1(把num[i]算上去)
		sum = int(math.Max(float64(itemMaxSum(nums, j)+1), float64(sum)))
	}
	return sum
}

// lengthOfLISDp
//  @Description: 使用动态规划改造
//  @param nums
//  @return int
func lengthOfLISDp(nums []int) int {
	maxSum := 0
	dp := make([]int, len(nums))

	for i, _ := range nums {
		dp[i] = 1
		for j := 0; j <= i; j++ {
			if nums[i] <= nums[j] {
				continue
			}
			dp[i] = int(math.Max(float64(dp[j]+1), float64(dp[i])))
		}
		maxSum = int(math.Max(float64(dp[i]), float64(maxSum)))
	}
	return maxSum
}
