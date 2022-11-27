package practice

import List "imooc-product/backend/链表"

/**
定义：数组中累积和与最小值的乘积，假设叫做指标A。给定一个数组，请返回子数组中，指标A最大的值。
 */
//实际就是以当前的数字为最小值, 找左边离他最近的比它小的数, 和右边最近比它小的数

//[5 3 2 1 6 7 8 4]
func QuotaA(nums []int) (int, map[string]int) {
	var stack []int //构建一个由小到大的单调栈
	max := -1
	maxMap := make(map[string]int)
	for i := 0; i < len(nums); i ++ {
		//遇到不符合大小的弹出栈, 直至符合条件
		for len(stack) != 0 && nums[stack[len(stack) - 1]] > nums[i] {
			index := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			end := i - 1
			start := 0
			if len(stack) != 0 {
				start = stack[len(stack) - 1] + 1
			}

			sum := 0
			for j := start; j <= end; j ++ {
				sum = sum + nums[j]
			}
			if max < sum * nums[index] {
				max = sum * nums[index]
				maxMap["start"] = start
				maxMap["end"] = end
			}
		}
		//压入栈
		stack = append(stack, i)
	}

	//清算阶段
	for i := len(stack) - 1; i >= 0; i -- {
		start := 0
		if i >= 1 {
			start = stack[i - 1] + 1
		}
		end := len(stack) - 1

		sum := 0
		for j := start; j <= end; j ++ {
			sum = sum + nums[j]
		}
		if max < sum * nums[i] {
			max = sum * nums[i]
			maxMap["start"] = start
			maxMap["end"] = end
		}
	}
	return max, maxMap
}

/**
输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
输出: [-1,3,-1]
解释:
    对于 num1 中的数字 4 ，你无法在第二个数组中找到下一个更大的数字，因此输出 -1 。
    对于 num1 中的数字 1 ，第二个数组中数字1右边的下一个较大数字是 3 。
    对于 num1 中的数字 2 ，第二个数组中没有下一个更大的数字，因此输出 -1 。
 */
func NextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, len(nums2))
	mapArr := make(map[int]int, len(nums1))
	var ret []int
	for i := 0; i < len(nums2); i++ {
		mapArr[nums2[i]] = -1
		for len(stack) > 0 && nums2[stack[len(stack) - 1]] < nums2[i] {
			mapArr[nums2[stack[len(stack) - 1]]] = nums2[i]
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}
	for _, v := range nums1 {
		ret = append(ret, mapArr[v])
	}
	return ret
}

func addTwoNumbers(l1 *List.Node, l2 *List.Node) *List.Node {
	var l3 *List.Node
	j := l3
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val
		if val >= 10 {
			val = (l1.Val + l2.Val) % 10
			if l1.Next != nil {
				l1.Next.Val = l1.Next.Val + 1
			}else {
				l1.Next = &List.Node{Val: 1, Next: nil}
			}
		}
		j = &List.Node{Val: val, Next: nil}
		j = j.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	return l3
}