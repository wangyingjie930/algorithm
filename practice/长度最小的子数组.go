package practice

import (
	"math"
)

/**
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
 */

func MinSubArrayLen(target int, nums []int) int {
	ret := processMinSubArrayLen(target, nums, 0)
	if ret == math.MaxInt32 {
		return -1
	}
	return ret
}

func processMinSubArrayLen(target int, nums []int, i int) int  {
	if i == len(nums) || target < 0 {
		if target == 0 {
			return 0
		}
		return -1
	}
	use := processMinSubArrayLen(target - nums[i], nums, i + 1)
	min := math.MaxInt32
	if use >= 0 {
		min = use + 1
	}
	nouse := processMinSubArrayLen(target, nums, i + 1)
	if nouse >= 0 && min > nouse {
		min = nouse
	}
	return min
}







/**
改造成记忆化搜索
实际就是添加一个数组, 缓存递归函数的可变参数对应的结果
而不用像暴力算法一样重新计算相同参数时的结果
 */
func MinSubArrayLenDp(target int, nums []int) int {
	var dp [][]int
 	for i := 0; i < len(nums) + 1; i ++ {
 		tmp := make([]int, target + 1)
		for j := 0; j <= target; j ++ {
			tmp[j] = -2
		}
		dp = append(dp, tmp)
	}
	ret := processMinSubArrayLenDp(target, nums, 0, dp)
	if ret == math.MaxInt32 {
		return -1
	}
	return ret
}

func processMinSubArrayLenDp(target int, nums []int, i int, dp [][]int) int  {
	if i == len(nums) || target < 0 {
		if target == 0 {
			dp[i][target] = 0
			return 0
		}
		return -1
	}

	if dp[i][target] != -2 {
		return dp[i][target]
	}

	use := processMinSubArrayLenDp(target - nums[i], nums, i + 1, dp)
	min := math.MaxInt32
	if use >= 0 {
		min = use + 1
	}
	nouse := processMinSubArrayLenDp(target, nums, i + 1, dp)
	if nouse >= 0 && min > nouse {
		min = nouse
	}
	dp[i][target] = min
	return min
}







/**
由递归关系可以推算出动态规划
使用动态规划-严格表结构
 */
func MinSubArrayLenDynamic(target int, nums []int) int {
	var dp [][]int
	//1. 确定可变参数, 2个, 那就是二维数组
	for i := 0; i <= len(nums); i ++ {
		tmp := make([]int, target + 1)
		for j := 0; j <= target; j ++ {
			if i == len(nums) {
				//2. 确定终止位置
				tmp[j] = -1
			}
		}
		//3. 确定basecase
		tmp[0] = 0
		dp = append(dp, tmp)
	}

	//4. 对于普遍的位置, 看依赖与其他的位置
	//5. 确定依赖的顺序
	//因为他是依赖于下方的位置, 那么就从数组底部开始遍历
	/**
	use := processMinSubArrayLenDp(target - nums[i], nums, i + 1, dp)
	min := math.MaxInt32
	if use >= 0 {
		min = use + 1
	}
	nouse := processMinSubArrayLenDp(target, nums, i + 1, dp)
	if nouse >= 0 && min > nouse {
		min = nouse
	}
	dp[i][target] = min
	 */
	for i := len(nums) - 1; i >= 0; i -- {
		for j := 1; j <= target; j ++ {
			nouse := dp[i + 1][j]
			use := -1
			if j-nums[i] >= 0 {
				use = dp[i + 1][j - nums[i]]
			}
			if nouse == -1 && use == -1 {
				dp[i][j] = -1
			}else if nouse == -1 || (use != -1 && use + 1 < nouse){
				dp[i][j] = use + 1
			}else {
				dp[i][j] = nouse
			}
		}
	}

	return dp[0][target]
}

//还有3维数组
//两个递归相互交叉
//

/**
2维数组+枚举
for (int index =N-1;index>=0;index--){
	for (int rest = e;rest <=aim;rest++){
		int ways =0;
		for (int zhang = O; arr[index]* zhang <= rest; zhang++) {
			ways += dp[index + 1] [rest - arr[index] * zhang];
		}
		dp[index][rest] = ways;
	}
}

根据位置依赖 将枚举优化
for (int index =N-1;index>=0;index--){
	for (int rest = e;rest <=aim;rest++){
		dp[index][rest] = dp[index+1][rest];
		if(rest - arr[index]>=0){
			dp[index][rest] += dp[index][rest - arr[index]];
		}
	}
}
 */

/**
如何进行尝试呢?
1. 尽量让可变参数是0维(int, string等, 不要数组)
2. 可变的参数尽可能地少(涉及到数组的维度)

 */

//滑动窗口
//打表法
//预处理法, 生成辅助数组, 加速答案的获取, 空间换时间
