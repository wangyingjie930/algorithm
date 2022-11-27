package practice

import (
	"math"
)

func MinWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	//当前窗口还需要多少字符才能满足条件
	//0 表示不需要了
	//负数表示重复了
	need := make(map[string]int)
	for _, v := range t {
		need[string(v)] ++
	}

	start, end := 0, 0
	minStart, minEnd := 0, math.MaxInt32
	for end < len(s) {
		val := string(s[end])
		if _, found := need[val]; found {
			need[val] --
		}
		for isEnough(need) && start <= end {
			if minEnd - minStart + 1 > end - start + 1 {
				minEnd, minStart = end, start
			}
			if _, found := need[string(s[start])]; found {
				need[string(s[start])] ++
			}
			start ++
		}
		end ++
	}
	if minEnd == math.MaxInt32 {
		 return ""
	}
	return s[minStart:minEnd + 1]
}

func isEnough(need map[string]int) bool {
	for _, v := range need{
		if v > 0 {
			return false
		}
	}
	return true
}
