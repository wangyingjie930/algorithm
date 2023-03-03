/**
  @author: wangyingjie
  @since: 2023/3/3
  @desc: https://leetcode.cn/problems/next-permutation/description/
  todo: 这个算是背起来吧
**/

package 下一个排列

// nextPermutation
//  @Description: 核心: 两次遍历交换逆序对
//  1. 尽可能靠右的低位进行交换
//	2. 将一个尽可能小的「大数」与前面的「小数」交换
//	3. 将「大数」换到前面后，需要将「大数」后面的所有数重置为升序，升序排列就是最小的排列
//  @param nums
func nextPermutation(nums []int) {
	//从后往前找出第一个出现递增的区间
	i, j := len(nums)-2, len(nums)-1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	if i >= 0 {
		//若找不到i则表示不存在一个字典序更大的排列
		//此时i处于以j为分界的左区间的"小数", 在j的右区间查找比nums[i]大的nums[k]
		k := len(nums) - 1
		//注意点: nums[k] <= nums[i]
		for k > j && nums[k] <= nums[i] {
			k--
		}
		//交换两数
		nums[i], nums[k] = nums[k], nums[i]
	}

	//从j开始改为递增
	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
