/**
  @author: wangyingjie
  @since: 2023/2/8
  @desc: 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
  @see:https://leetcode.cn/problems/maximum-subarray/description/
**/

package 最大子数组和

import "math"

// maxSubArray
//  @Description: 递归版本
//  @param nums
//  @return int
func maxSubArray(nums []int) int {
	sum := math.MinInt8
	for i := 0; i < len(nums); i++ {
		// 求出每个以i为结束节点的最大子序列之和
		// 再在这里面中选出最大的
		sum = int(math.Max(float64(maxSumByBackItem(nums, i)), float64(sum)))
	}
	return sum
}

// maxSumByBackItem
//  @Description: 求出以i节点结束的子序列中最大的长度是多少
//  @param nums
//  @param i
//  @return int
func maxSumByBackItem(nums []int, i int) int {
	if i == 0 {
		return nums[0]
	}
	return int(math.Max(float64(maxSumByBackItem(nums, i-1)+nums[i]), float64(nums[i])))
}

func maxSubArrayDynamic(nums []int) int {
	sum := math.MinInt64
	plan := nums
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			plan[i] = nums[i]
		} else {
			plan[i] = int(math.Max(float64(plan[i-1]+nums[i]), float64(nums[i])))
		}
		sum = int(math.Max(float64(plan[i]), float64(sum)))
	}
	return sum
}
