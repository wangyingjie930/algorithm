/**
  @author: wangyingjie
  @since: 2023/3/26
  @desc: https://leetcode.cn/problems/combination-sum/description/
**/

package 组合总和

func combinationSum(candidates []int, target int) [][]int {
	var cur []int
	var ret [][]int
	helper(cur, 0, candidates, target, &ret)
	return ret
}

// helper
//  @Description: 回溯算法最基本的要有 当前选择路径, 剩余选择, 存放结果集
//  @param cur
//  @param start
//  @param candidates
//  @param target
//  @param ret
func helper(cur []int, start int, candidates []int, target int, ret *[][]int) {
	if target == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*ret = append(*ret, tmp)
		return
	} else if target < 0 {
		return
	}

	//这里的start要重点注意
	//题目要求包含重复的那就要固定为0
	//题目要求不包含重复的就要传参, 表示去重,不回头
	for i := start; i < len(candidates); i++ {
		cur = append(cur, candidates[i])
		helper(cur, i, candidates, target-candidates[i], ret)
		cur = cur[:len(cur)-1]
	}
}
