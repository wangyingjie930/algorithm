/**
  @author: wangyingjie
  @since: 2023/4/2
  @desc: https://leetcode.cn/problems/maximum-product-subarray/
**/

package 乘积最大子数组

import "math"

// maxProduct
//  @Description: 重点: 不能单单dp[i] = max(dp[i], dp[i-1]*nums[i]), 因为这样考虑不到负负得正的情况
//  @param nums
//  @return int
func maxProduct(nums []int) int {
	dp := make([][]int, len(nums))
	maxSum := math.MinInt64
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, 2)
		if nums[i] > 0 {
			dp[i][1] = nums[i]
			maxSum = max(maxSum, dp[i][1])
		} else {
			dp[i][0] = nums[i]
			maxSum = max(maxSum, dp[i][0])
		}
	}

	for i := 1; i < len(nums); i++ {
		//dp[i][0]表示以i为底的连续数组乘积的最小值(负数)是什么
		dp[i][0] = min(dp[i][0], dp[i-1][0]*nums[i])
		dp[i][0] = min(dp[i][0], dp[i-1][1]*nums[i])

		//dp[i][1]表示以i为底的连续数组乘积的最大值(正数)是什么
		dp[i][1] = max(dp[i][1], dp[i-1][1]*nums[i])
		dp[i][1] = max(dp[i][1], dp[i-1][0]*nums[i])

		maxSum = max(max(maxSum, dp[i][0]), dp[i][1])
	}

	return maxSum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
