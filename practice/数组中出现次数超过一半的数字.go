/**
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
 */
package practice

import "fmt"

/**
使用选择排序实现, 超出时间限制
 */
func MajorityElementSelectSort(nums []int) int {
	nums = selectSort(nums)
	fmt.Println(nums)
	return nums[len(nums) / 2]
}

func selectSort(arr []int) []int  {
	for i := 0; i < len(arr) / 2 + 1; i++ {
		min := arr[i]
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}


/**
使用归并的方法
 */
func MajorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	leftMain := MajorityElement(nums[0:len(nums) / 2])
	rightMain := MajorityElement(nums[len(nums) / 2:])

	if leftMain == rightMain {
		 return leftMain
	}
	if leftCountBigger(nums, leftMain, rightMain) {
		return leftMain
	}
	return rightMain
}

/**
左边数组最多数字的数量是否大于右边数组最多数字的数量
 */
func leftCountBigger(nums []int, left int, right int) bool {
	leftCount := 0
	rightCount := 0
	for _, v := range nums {
		if v == left {
			leftCount ++
		}
		if v == right {
			rightCount ++
		}
	}
	return leftCount > rightCount
}
