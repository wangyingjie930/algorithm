/**
  @author: wangyingjie
  @since: 2023/2/22
  @desc: https://leetcode.cn/problems/trapping-rain-water/
**/

package 接雨水_遍历柱子

import "math"

// trap
//  @Description: 按柱子遍历
//  @param height
//  @return int
func trap(height []int) int {
	res := 0
	for i := 1; i < len(height)-1; i++ {
		//求出每个柱子左右两边大于该柱子的边界
		left := i
		for l := i - 1; l >= 0; l-- {
			if height[l] > height[i] && height[l] > height[left] {
				left = l
			}
		}

		right := i
		for r := i + 1; r < len(height); r++ {
			if height[r] > height[i] && height[r] > height[right] {
				right = r
			}
		}
		if left != i && right != i {
			//只需求出i这根柱子距离最低水位还差多少格子即可
			res = res + int(math.Min(float64(height[left]), float64(height[right]))) - height[i]
		}
	}
	return res
}

// trapDp
//  @Description: 动态规划
//  @param height
//  @return int
func trapDp(height []int) int {
	res := 0

	//i柱子左右两边分别的最大值是什么
	maxLeft := make([]int, len(height))
	for i := 1; i < len(height); i++ {
		maxLeft[i] = int(math.Max(float64(maxLeft[i-1]), float64(height[i-1])))
	}
	maxRight := make([]int, len(height))
	for i := len(height) - 2; i > 0; i-- {
		maxRight[i] = int(math.Max(float64(maxRight[i+1]), float64(height[i+1])))
	}

	for i := 1; i < len(height)-1; i++ {
		if maxLeft[i] > height[i] && maxRight[i] > height[i] {
			//只需求出i这根柱子距离最低水位还差多少格子即可
			res = res + int(math.Min(float64(maxLeft[i]), float64(maxRight[i]))) - height[i]
		}
	}
	return res
}
