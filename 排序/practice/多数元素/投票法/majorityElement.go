/**
  @author: wangyingjie
  @since: 2023/4/3
  @desc: https://leetcode.cn/problems/majority-element/solutions/?orderBy=most_votes
**/

package 投票法

// majorityElement
//  @Description: 最坏的情况下非众数中的每一个数都和众数进行消除，那么剩下的是众数。其他情况则显然剩下的也是众数本身
//  @param nums
//  @return int
func majorityElement(nums []int) int {
	count := 1
	target := nums[0]

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			target = nums[i]
			count = 1
		} else if target == nums[i] {
			count++
		} else {
			count--
		}
	}
	return target
}
