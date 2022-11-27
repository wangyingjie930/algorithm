package practice

import (
	"math"
	"strings"
)

func SolveNQueens(n int) [][]string {
	var ret [][]string
	record := make([]int, n)
	process(0, record, n, &ret)
	return ret
}

/**
record = [1, 2, 3...]表示
[
	第0行的皇后 => 第1列,
	第1行的皇后 => 第2列,.....
]
record变量的含义是0到i-1行已经符合规定的相应位置
 */
func process(i int, record []int, n int, ret *[][]string)  {
	if i == n {
		//表明已经到底了, 可以进行输出了
		*ret = append(*ret, printArr(record))
		return
	}
	for j := 0; j < n; j++ {
		if isValid(record, i, j) {
			//第i行存放在第j列的时候是有效的
			record[i] = j
			//往下继续查找
			process(i + 1, record, n, ret)
		}
	}
	return
}

func isValid(record []int, i, j int) bool {
	for k := 0; k < i; k ++ {
		//遍历之前的皇后, 规则: 不在同一列和同一斜线
		if j == record[k] ||
			math.Abs(float64(j) - float64(record[k])) == math.Abs(float64(i) - float64(k)) {
			return false
		}
	}
	return true
}

func printArr(record []int) []string {
	var ret []string
	for _, v := range record {
		base := []byte(strings.Repeat(".", len(record)))
		base[v] = 'Q'
		ret = append(ret, string(base))
	}
	return ret
}