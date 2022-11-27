package practice

import (
	"math"
	"strconv"
)

func NumDecodings(s string) int {
	bytes := []byte(s)
	if s == "0" || (len(s) > 0 && string(bytes[0]) == "0") {
		return 0
	}
	if len(s) <= 1 {
		return 1
	}
	num := 0
	if len(s) > 1 {
		num = NumDecodings(string(bytes[1:]))

		two, _ := strconv.Atoi(string(bytes[:2]))
		if two <= 26 {
			num = num + NumDecodings(string(bytes[2:]))
		}
	}
	return num
}

func NumDecodings1(s string) int {
	return processNumDecodings(s, 0)
}

/**
使用动态规划的方法进行优化
 */
func processNumDecodings(s string, i int) int {
	if s == "0" || (len(s) > 0 && string(s[0]) == "0") {
		return 0
	}
	dp := make([]int, len(s) + 1)
	dp[len(s)] = 1
	if string(s[len(s) - 1]) != "0" {
		dp[len(s) - 1] = 1
	}
	for i := len(s) - 2; i >= 0; i-- {
		if string(s[i]) == "0" {
			continue
		}
		dp[i] = dp[i + 1]
		if two, err := strconv.Atoi(s[i : i+2]); err == nil && two <= 26 {
			dp[i] = dp[i] + dp[i + 2]
		}
	}
	return dp[i]
}

/*
啤酒 2 块钱一瓶，4 个盖换一瓶，2 个空瓶换一瓶。
问:10 块钱可以喝几瓶?
*/

func Beer1(money int) int {
	if money == 2 {
		return 1
	}
	if money < 2 {
		return 0
	}
	max := math.MinInt32
	i := 1
	for money > 2 {
		tmp := Beer1(money - i * 2) + i + i / 4 + i / 2
		if max < tmp {
			max = tmp
		}
		money = money - i * 2
		i ++
	}
	return max
}