/**
  @author: wangyingjie
  @since: 2023/3/25
  @desc:https://leetcode.cn/problems/coin-change/
**/

package 零钱兑换

// coinChange
//  @Description:
/**
状态: 金额数量
选择: 不同面值的硬币
dp[状态] = 选择的不同方案
*/
//  @param coins
//  @param amount
//  @return int
func coinChange(coins []int, amount int) int {
	arr := make([]int, amount+1) //这里+1是因为要用下标来表示金额, 单纯amount表示不了arr[amount]
	for i := 0; i <= amount; i++ {
		//这里是给个默认值, 因为后面要进行的是去最小值, 需要取个足够大的数
		arr[i] = amount + 1
	}

	arr[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i < coins[j] {
				//表示肯定凑不到
				continue
			}
			//point
			arr[i] = min(arr[i], arr[i-coins[j]]+1)
		}
	}

	if arr[amount] == amount+1 {
		return -1
	}
	return arr[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
