/**
  @author: wangyingjie
  @since: 2023/3/22
  @desc: https://leetcode.cn/problems/string-to-integer-atoi/
**/

package 字符串转换整数

import "math"

func myAtoi(s string) int {
	startInt := false
	ret := 0
	isNeg := false
	for i := 0; i < len(s); i++ {
		if !startInt {
			if s[i] == '-' || s[i] == '+' || s[i] == ' ' || (s[i] >= '0' && s[i] <= '9') {
				if s[i] == ' ' {
					continue
				}
				startInt = true
				if s[i] == '-' {
					isNeg = true
					continue
				}
				if s[i] == '+' {
					continue
				}
			} else {
				break
			}
		}

		if startInt {
			if s[i] >= '0' && s[i] <= '9' {
				ret = ret*10 + int(s[i]-'0')
				if ret > math.MaxInt32 {
					break
				}
			} else {
				break
			}
		}
	}

	if isNeg {
		return int(math.Max(float64(ret*-1), math.MinInt32))
	}
	return int(math.Min(float64(ret), math.MaxInt32))
}
