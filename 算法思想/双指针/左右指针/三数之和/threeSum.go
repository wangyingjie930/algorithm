/**
  @author: wangyingjie
  @since: 2023/3/2
  @desc: https://leetcode.cn/problems/3sum/
**/

package 三数之和

import "sort"

func threeSum(nums []int) [][]int {
	//sortArray(nums)
	//先进行排序
	sort.Ints(nums)
	var ret [][]int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] {
			//注意去重, 越过相同的数
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			leftItem, rightItem := nums[left], nums[right]
			//三数相加
			sum := nums[i] + leftItem + rightItem
			if sum == 0 {
				ret = append(ret, []int{nums[i], leftItem, rightItem})
				for left < right && nums[right] == rightItem {
					//右指针左移, 注意去重, 右指针越过相同的数
					right--
				}
				for left < right && nums[left] == leftItem {
					//左指针右移, 注意去重, 左指针越过相同的数
					left++
				}
			} else if sum > 0 {
				for left < right && nums[right] == rightItem {
					//注意去重, 右指针越过相同的数
					right--
				}
			} else {
				for left < right && nums[left] == leftItem {
					//注意去重, 左指针越过相同的数
					left++
				}
			}
		}
	}
	return ret
}

// sortArray
//  @Description: 双指针前提, 有序数组
//  @param nums
func sortArray(nums []int) {
	for i := 0; i < len(nums); i++ {
		val := nums[i]
		for j := i - 1; j >= 0; j-- {
			if nums[j] > val {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
