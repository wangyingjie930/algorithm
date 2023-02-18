/**
  @author: wangyingjie
  @since: 2023/2/14
  @desc: https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/
**/

package 将有序数组转换为二叉搜索树

import (
	"algorithm/二叉查找树/AVL"
)

func sortedArrayToBST(nums []int) *AVL.Node {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, start, end int) *AVL.Node {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	node := &AVL.Node{Value: nums[mid], Key: nums[mid]}
	node.Left = helper(nums, start, mid-1)
	node.Right = helper(nums, mid+1, end)
	return node
}
