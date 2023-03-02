/**
  @author: wangyingjie
  @since: 2023/3/2
  @desc: https://leetcode.cn/problems/path-sum-ii/
**/

package 路径总和II

import "algorithm/二叉查找树/Tree"

func pathSum(root *Tree.Node, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	cur := []*Tree.Node{root}
	var ret [][]int
	//注意点: 这里的cur和target都要先扣除根节点
	helper(root, cur, targetSum-root.Value, &ret)
	return ret
}

// helper
//  @Description: 回溯需要的变量: 已选择的路径, 待选择的列表(因为是树可忽略), 全局变量
//  @param root
//  @param cur 已选择的路径
//  @param target
//  @param ret 全局变量
func helper(root *Tree.Node, cur []*Tree.Node, target int, ret *[][]int) {
	if root.Left == nil && root.Right == nil {
		if target == 0 {
			*ret = append(*ret, make([]int, len(cur)))
			for i, node := range cur {
				(*ret)[len(*ret)-1][i] = node.Value
			}
		}
		return
	}

	if root.Left != nil {
		cur = append(cur, root.Left)
		//注意点: 记得传左节点, 还有target要扣减左节点的值
		helper(root.Left, cur, target-root.Left.Value, ret)
		cur = cur[0 : len(cur)-1]
	}
	if root.Right != nil {
		cur = append(cur, root.Right)
		helper(root.Right, cur, target-root.Right.Value, ret)
		cur = cur[0 : len(cur)-1]
	}
}
