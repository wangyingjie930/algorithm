/**
  @author: wangyingjie
  @since: 2023/2/25
  @desc: https://leetcode.cn/problems/edit-distance/
**/

package 编辑距离

import "math"

func minDistance(word1 string, word2 string) int {
	return dp(word1, word2)
}

func helper(source string, target string, sourceIndex, targetIndex int) int {
	if sourceIndex >= len(source) {
		return len(target) - targetIndex
	}
	if targetIndex >= len(target) {
		/**
		  aaaaaaaa
		bb
		假设这种情况: 一直选择插入, 导致target(bb)已经到达底部了, sourceIndex还是指向source[0], 那么需要减掉aaaaaaaa的长度
		*/
		return len(source) - sourceIndex
	}
	if source[sourceIndex] == target[targetIndex] {
		return helper(source, target, sourceIndex+1, targetIndex+1)
	} else {
		//新增
		min := float64(helper(source, target, sourceIndex, targetIndex+1) + 1)

		//删除
		del := float64(helper(source, target, sourceIndex+1, targetIndex) + 1)
		if min > del {
			min = del
		}

		//更新
		upd := float64(helper(source, target, sourceIndex+1, targetIndex+1) + 1)
		if min > upd {
			min = upd
		}
		return int(min)
	}
}

func dp(source, target string) int {
	table := make([][]int, len(target)+1)
	for i := 0; i <= len(target); i++ {
		table[i] = make([]int, len(source)+1)
		table[i][len(source)] = len(target) - i
	}
	for i := 0; i <= len(source); i++ {
		table[len(target)][i] = len(source) - i
	}
	for i := len(target) - 1; i >= 0; i-- {
		for j := len(source) - 1; j >= 0; j-- {
			if source[j] == target[i] {
				table[i][j] = table[i+1][j+1]
			} else {
				table[i][j] =
					int(math.Min(
						math.Min(float64(table[i+1][j])+1, float64(table[i][j+1])+1),
						float64(table[i+1][j+1])+1,
					))
			}
		}
	}
	return table[0][0]
}
