/**
  @author: wangyingjie
  @since: 2023/2/16
  @desc: https://leetcode.cn/problems/first-missing-positive/description/
  @see : firstMissingPositive.png
**/

package 缺失的第一个正数

import "math"

// firstMissingPositive
//  @Description: 已知数组长度为n, 取正数肯定在[1, n+1]这个范围中拿
//  @param nums
//  @return int
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			nums[i] = len(nums) + 100 //这里取的是一个不会不会出现的值 (因为是负数, 所以找不到位置标记, 等价与取超出数组长度N的数)
		}
	}

	for i := 0; i < len(nums); i++ {
		num := abs(nums[i]) //这里是防止被标记过后又被使用, 为了索引下标先转回来
		if num-1 >= 0 && num-1 < len(nums) {
			//实际和哈希表 <num, 1>的原理一样, 只是这里的num-1是下标, *-1是标记
			nums[num-1] = abs(nums[num-1]) * -1
		}
	}

	i := 1
	for ; i <= len(nums); i++ {
		if nums[i-1] > 0 { //这里就是判断没有被标记时, 返回
			return i
		}
	}
	return i
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}
