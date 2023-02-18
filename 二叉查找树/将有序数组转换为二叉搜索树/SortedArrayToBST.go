/**
  @author: wangyingjie
  @since: 2023/2/14
  @desc: https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/
**/

package 将有序数组转换为二叉搜索树

import (
	"algorithm/二叉查找树/Tree"
)

func sortedArrayToBST(nums []int) *Tree.Node {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, start, end int) *Tree.Node {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	node := &Tree.Node{Value: nums[mid], Key: nums[mid]}
	node.Left = helper(nums, start, mid-1)
	node.Right = helper(nums, mid+1, end)
	return node
}
