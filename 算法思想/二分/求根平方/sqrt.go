package 求根平方

import (
	"fmt"
	"strconv"
)

func mySqrt(x int) int {
	l, r := 0, x
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l - 1
}

// mySqrtAccuracy
//  @Description: 保留小数点n位
//  @param x
//  @param n
//  @return float64
func mySqrtAccuracy(x int, n int) float64 {
	l, r := float64(0), float64(x)
	e := 1e-10 //这里表示精确的误差为10的-10方
	for l <= r {
		mid := (l + r) / 2
		if mid <= float64(x)/mid {
			l = mid + e
		} else {
			r = mid - e
		}
	}

	//将求出的结果保留小数点n位
	str := "%." + strconv.Itoa(n) + "f"
	ret, _ := strconv.ParseFloat(fmt.Sprintf(str, l-e), 64)
	return ret
}
