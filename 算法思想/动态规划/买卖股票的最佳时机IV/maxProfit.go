/**
  @author: wangyingjie
  @since: 2023/2/27
  @desc:
	https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/
	https://labuladong.github.io/algo/di-ling-zh-bfe1b/yi-ge-fang-3b01b/
**/

package 买卖股票的最佳时机IV

import (
	"math"
)

func maxProfit(k int, prices []int) int {
	//return helper(prices, k, len(prices)-1, 0)
	return dpFunc(prices, k)
}

// helper
// 分析:
// [选择] 买入, 卖出, 保持不变
// [状态] 天数变动, 交易次数限制, 持有股票限制
//
//  @Description: 返回持有的利润
//  @param prices
//  @param k 限制的买入次数
//  @param day 天数
//  @param keep 当前是否持有
//  @return int
func helper(prices []int, k int, day int, keep int) int {
	if k == 0 && keep == 1 {
		//没有买入次数, 但却拥有股票, 因为后面是最大值对比, 所以给个最小值
		return math.MinInt64
	}
	if k == 0 && keep == 0 {
		//没有买入次数, 没有股票, 返回0利润
		return 0
	}
	if day == -1 && keep == 0 {
		//为什么day选-1? 因为当day=0时可以选择买入(此时的利润为负的), 不能确定默认值
		//当day=-1, 且当前没有股票, 返回0利润
		return 0
	}
	if day == -1 && keep == 1 {
		//当day=-1, 但却拥有股票, 因为后面是最大值对比, 所以给个最小值
		return math.MinInt64
	}

	if keep == 0 {
		//若当前并不持有股票
		return int(math.Max(
			//可能是 前一天(day-1)并没有持有(keep=0)股票, 因为是保持, 不用为当前腾出一次操作的机会(k)
			float64(helper(prices, k, day-1, 0)),
			//可能是 前一天(day-1)持有了股票(keep=1)股票, 并在今天进行股票的卖出(+prices[day]),
			//因为是卖出, 不用为当前腾出一次操作的机会(k)
			float64(helper(prices, k, day-1, 1)+prices[day]),
		))
	} else if keep == 1 {
		//若当前持有股票
		return int(math.Max(
			//可能是 前一天(day-1)并持有(keep=1)股票, 因为是保持, 不用为当前腾出一次操作的机会(k=0)
			float64(helper(prices, k, day-1, 1)),
			//可能是 前一天(day-1)不持有有了股票(keep=0)股票, 并在今天进行股票的卖入(-prices[day]),
			//因为是买入, 需要将当前的一次操作记录, 剩余k-1次限制留给子问题
			float64(helper(prices, k-1, day-1, 0)-prices[day]),
		))
	}
	return 0
}

// dpFunc
// [选择] 买入, 卖出, 保持不变
// [状态] 天数变动, 交易次数限制, 持有股票限制
// 动态规划算法本质上就是穷举「状态」，然后在「选择」中选择最优解
//
/**
for 状态1 in 状态1的所有取值：
    for 状态2 in 状态2的所有取值：
        for ...
            dp[状态1][状态2][...] = 择优(选择1，选择2...)

*/
//  @Description:
//  @param prices
//  @param k
//  @return int
func dpFunc(prices []int, k int) int {
	dp := make([][][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([][]int, len(prices)+1)
		for j := 0; j < len(prices)+1; j++ {
			dp[i][j] = make([]int, 2)
			dp[0][j][0] = 0
			dp[0][j][1] = math.MinInt64
			dp[i][0][0] = 0
			dp[i][0][1] = math.MinInt64

			if i > 0 && j > 0 {
				dp[i][j][0] = int(math.Max(float64(dp[i][j-1][0]), float64(dp[i][j-1][1]+prices[j-1])))
				dp[i][j][1] = int(math.Max(float64(dp[i][j-1][1]), float64(dp[i-1][j-1][0]-prices[j-1])))
			}
		}
	}

	return dp[k][len(prices)][0]
}
