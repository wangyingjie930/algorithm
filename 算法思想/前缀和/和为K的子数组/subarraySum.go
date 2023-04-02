/**
  @author: wangyingjie
  @since: 2023/4/2
  @desc: https://leetcode.cn/problems/subarray-sum-equals-k/
**/

package 和为K的子数组

func subarraySum(nums []int, k int) int {
	mapper := make(map[int]int, len(nums))
	mapper[0] = 1
	pre := 0
	count := 0

	for i := 0; i < len(nums); i++ {
		pre = pre + nums[i]

		//开始处理索引以i为结尾的, 总和为k的连续数组
		//需要先处理, 以确保mapper里面维护的是[0, i-1]的前缀和
		if _, exist := mapper[pre-k]; exist {
			count += mapper[pre-k]
		}

		//维护i的前缀和映射表
		if _, exist := mapper[pre]; exist {
			mapper[pre] += 1
		} else {
			mapper[pre] = 1
		}
	}

	return count
}
