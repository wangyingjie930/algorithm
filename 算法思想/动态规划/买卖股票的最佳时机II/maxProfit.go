/**
  @author: wangyingjie
  @since: 2023/4/2
  @desc: https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
**/

package 买卖股票的最佳时机II

func maxProfit(prices []int) int {
	//dp[i][0]: 表示直到第i天不持有股票, dp[i][1]: 表示直到第i天还持有股票

	//dp[i][0] = dp[i-1][1] + prices[i], dp[i-1][0]: 买入或者保持不动
	//dp[i][1] = dp[i-1][0] - prices[i], dp[i-1][1]: 卖出或者保持不动
	//max(dp[i][0], dp[i][1])
	dp := make([][]int, len(prices))
	maxPro := 0
	dp[0] = []int{0: 0, 1: 0 - prices[0]}

	for i := 1; i < len(prices); i++ {
		dp[i] = make([]int, 2)
		dp[i][0] = max(dp[i-1][1]+prices[i], dp[i-1][0])
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
		maxPro = max(max(maxPro, dp[i][0]), dp[i][1])
	}

	return maxPro
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
