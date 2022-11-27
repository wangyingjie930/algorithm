package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 0, 0, 0}
	merge1(a, 3, []int{2, 5, 6}, 3)
	fmt.Printf("%+v", a)
}

//使用额外的空间进行
func merge(nums1 []int, m int, nums2 []int, n int)  {
	var ret []int
	i, j := 0, 0
	for ; i < m && j < n; {
		if nums1[i] < nums2[j] {
			ret = append(ret, nums1[i])
			i ++
		}else {
			ret = append(ret, nums2[j])
			j ++
		}
	}

	if i >= m {
		ret = append(ret, nums2[j:]...)
	}
	if j >= n {
		ret = append(ret, nums1[i:]...)
	}
	copy(nums1, ret)
}

//从后往前开始遍历, 在num1上不使用额外的数组
func merge1(nums1 []int, m int, nums2 []int, n int)  {
	i, j := m - 1, n - 1
	k := len(nums1) - 1
	for ; i >= 0 || j >= 0; {
		if i < 0 {
			nums1[k] = nums2[j]
			j --
		}else if j < 0 {
			nums1[k] = nums1[i]
			i --
		}else if nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			j --
		}else {
			nums1[k] = nums1[i]
			i --
		}
		k --
	}
}