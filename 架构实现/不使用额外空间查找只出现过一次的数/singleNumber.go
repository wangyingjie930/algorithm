/**
  @author: wangyingjie
  @since: 2023/4/11
  @desc: 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素
**/

package singleNumber

// singleNumber
//  @Description: 异或:一个运算子为true，另一个为false 才会返回 true
// 一个数异或它本身2次, 还是等于它本身
//  @param nums
//  @return int
func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}
