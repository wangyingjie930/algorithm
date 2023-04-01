/**
  @author: wangyingjie
  @since: 2023/4/2
  @desc: https://leetcode.cn/problems/house-robber/
**/

package 打家劫舍

// rob
//  @Description: dp[i][0] 表示第i家不偷, dp[i][1] 表示第i家要偷
//  @param nums
//  @return int
func rob(nums []int) int {
	dp := make([][]int, len(nums))
	maxRob := 0

	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, 2)
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = nums[i]
		} else {
			dp[i][0] = max(dp[i-1][1], dp[i-1][0]) //当第i家不偷的时候, 可以选择上一家的最大值
			dp[i][1] = dp[i-1][0] + nums[i]        //当第i家偷的时候, 必须让上一家不被偷
		}
		maxRob = max(maxRob, dp[i][0])
		maxRob = max(maxRob, dp[i][1])
	}

	return maxRob
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
