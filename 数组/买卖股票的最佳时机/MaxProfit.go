/**
  @author: wangyingjie
  @since: 2023/2/8
  @desc:和最大子数组和类似
  @see: MaxSubArray.go
  @see: https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/
**/

package 买卖股票的最佳时机

import "math"

// maxProfit
//  @Description: 仍旧是遍历每个元素, 以每个元素作为结束节点或者右边界去求解子问题, 最后通过汇总的形式挑选
//  @param prices
//  @return int
func maxProfit(prices []int) int {
	maxPro := 0
	for i := 0; i < len(prices); i++ {
		minItem := minBeforeItem(prices, i)
		maxPro = int(math.Max(float64(maxPro), float64(prices[i]-minItem)))
	}
	return maxPro
}

// minBeforeItem
//  @Description: 以i为右边界获取边界内最小的数
//  @param prices
//  @param i
//  @return item
func minBeforeItem(prices []int, i int) (item int) {
	if i == 0 {
		return prices[i]
	}

	ret := minBeforeItem(prices, i-1)
	if ret < prices[i] {
		return ret
	} else {
		return prices[i]
	}
}

// maxProfitDynamic
//  @Description: 动态规划 根据minBeforeItem做改造
//  @param prices
//  @return int
func maxProfitDynamic(prices []int) int {
	maxPro := 0
	plan := make([]int, len(prices))

	for i := 0; i < len(prices); i++ {
		if i == 0 {
			plan[i] = prices[i]
		} else {
			plan[i] = int(math.Min(float64(plan[i-1]), float64(prices[i])))
		}
		maxPro = int(math.Max(float64(maxPro), float64(prices[i]-plan[i])))
	}
	return maxPro
}
