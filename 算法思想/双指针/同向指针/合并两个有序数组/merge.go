/**
  @author: wangyingjie
  @since: 2023/3/18
  @desc: https://leetcode.cn/problems/merge-sorted-array/
  掌握程度: 原地合并思路需要题解辅助
**/

package 合并两个有序数组

// merge
//  @Description: 核心就是双指针从后往前遍历
//  @param nums1
//  @param m
//  @param nums2
//  @param n
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j, k := m-1, n-1, m+n-1; j >= 0; k-- {
		if i < 0 || nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
	}
}
