/**
  @author: wangyingjie
  @since: 2023/3/21
  @desc: https://leetcode.cn/problems/two-sum/description/
**/

package 两数之和

func twoSum(nums []int, target int) []int {
	mapper := make(map[int]int, len(nums))
	for i, num := range nums {
		if index, exist := mapper[num]; exist {
			//注意和循环放在一起减少1次遍历
			return []int{index, i}
		}
		mapper[target-num] = i
	}
	return []int{}
}
