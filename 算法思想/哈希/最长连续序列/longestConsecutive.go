/**
  @author: wangyingjie
  @since: 2023/3/28
  @desc: https://leetcode.cn/problems/longest-consecutive-sequence/
**/

package 最长连续序列

// longestConsecutive
//  @Description: 思路, 遍历数组判断是否可以以它开头, 并每次循环+1, 查找nums有没有+1后的数, 求最大长度
//  @param nums
//  @return int
func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool, len(nums))
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		numSet[nums[i]] = true
	}

	for i := 0; i < len(nums); i++ {
		//先观察数组里面是否有以当前的数开头的
		if _, exist := numSet[nums[i]-1]; exist {
			//如果存在则跳过
			continue
		}
		//如果不存在, 证明可以以它为开头, 按顺序查找, 并将顺序的长度记录
		length := 1
		for j := nums[i] + 1; true; {
			if _, exist := numSet[j]; !exist {
				break
			}

			j++
			length++
		}
		if maxLen < length {
			//更新最大的顺序长度
			maxLen = length
		}
	}

	return maxLen
}
