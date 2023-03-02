/**
  @author: wangyingjie
  @since: 2023/3/2
  @desc: https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
**/

package 买卖股票的最佳时机

import "math"

func maxProfit(prices []int) int {
	//状态: 天数, 是否持有
	//选择: 买入, 卖出, 持有

	dpTable := make([][]int, len(prices)+1)
	dpTable[0] = make([]int, 2)
	dpTable[0][0] = 0
	dpTable[0][1] = math.MinInt64

	for i := 1; i < len(prices)+1; i++ {
		dpTable[i] = make([]int, 2)
		//当前不持有股票时, 可能是昨天也是不持有, 或者昨天持有今天卖掉了
		dpTable[i][0] = int(math.Max(float64(dpTable[i-1][1]+prices[i-1]), float64(dpTable[i-1][0])))
		//dpTable[i][1] = int(math.Max(float64(dpTable[i-1][0]-prices[i-1]), float64(dpTable[i-1][1])))
		// 因为只能限制一次买入, 所以用的是0元去购买, 如果改成dpTable[i-1][0]很可能是在交易多次的基础上进行购买
		dpTable[i][1] = int(math.Max(float64(-prices[i-1]), float64(dpTable[i-1][1])))
	}

	return dpTable[len(prices)][0]
}
