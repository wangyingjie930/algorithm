/**
  @author: wangyingjie
  @since: 2023/2/25
  @desc: https://leetcode.cn/problems/n-queens/
**/

package N皇后

import "math"

func solveNQueens(n int) [][]string {
	var ret [][]string
	used := make(map[int]bool, n)
	cur := make([][]byte, n)
	for i := 0; i < n; i++ {
		cur[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			cur[i][j] = '.'
		}
	}
	dfs(0, n, used, cur, &ret)
	return ret
}

// dfs
//  @Description:
//  @param row 第几行
//  @param N 共几行几列
//  @param used 【选择列表】
//  @param cur 【选择路径】
//  @param ret 【结果】
func dfs(row int, N int, used map[int]bool, cur [][]byte, ret *[][]string) {
	if row == N {
		//将摆放的结果输出到ret
		var tmp []string
		for i := 0; i < len(cur); i++ {
			tmp = append(tmp, string(cur[i]))
		}
		*ret = append(*ret, tmp)
		return
	}

	//多叉树是根据摆放不同的列进行分裂节点的
	//这里循环的是列数（N）
	for i := 0; i < N; i++ {
		if !used[i] {
			//进行位置摆放的判断
			if !isValid(cur, row, i) {
				continue
			}
			used[i] = true
			cur[row][i] = 'Q'
			//往下一行开始
			dfs(row+1, N, used, cur, ret)
			cur[row][i] = '.'
			used[i] = false
		}
	}
}

// isValid
//  @Description: 判断该位置是否符合条件
//  @param cur
//  @param row
//  @param col
//  @return bool
func isValid(cur [][]byte, row int, col int) bool {
	//因为是从上倒下，从左到右遍历，仅需判断上，左上即可
	for i := 0; i < row; i++ {
		for j := 0; j < len(cur[i]); j++ {
			if cur[i][j] == 'Q' {
				if j == col || math.Abs(float64(col)-float64(j)) == math.Abs(float64(row)-float64(i)) {
					return false
				}
				break
			}
		}
	}
	return true
}
