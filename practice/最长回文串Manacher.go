package practice

import (
	"strings"
)

/**
manacher算法实现获取最长回文子串的长度
 */
func MaxLcpsLength(s string) int {
	var buffer []string
	for _, v := range s {
		buffer = append(buffer, string(v))
	}
	newstring := strings.Join(buffer, "#")

	c := -1
	r := -1
	bufferLen := make([]int, len(newstring))//以每个字符为中心的最长回文子串长度
	max := -1
	for k, _ := range newstring {
		bufferLen[k] = 1
		if k >= r {
			//第一种情况: 当前字符在R外面, 暴力扩
			for i := 1; k - i >= 0 && k + i < len(newstring) && newstring[k - i] == newstring[k + i]; i ++ {
				bufferLen[k] = bufferLen[k] + 2
				if k + i > r {
					r = k + i
					c = k
				}
			}
		}else {
			//第2种情况: 当前字符在R里面, 取k1为k的对称点, k的最小回文半径等于k1
			k1 := k - (k - c) * 2
			l := r - (r - c) * 2
			if k1 < 0 || l < 0 {
				continue
			}
			lk1 := k1 - bufferLen[k1] / 2
			bufferLen[k] = bufferLen[k1]
			if lk1 == l {
				for i := r - k; k - i >= 0 && k + i < len(newstring) && newstring[k - i] == newstring[k + i]; i ++ {
					bufferLen[k] = bufferLen[k] + 2
					if k + i > r {
						r = k + i
						c = k
					}
				}
			}
		}
		if max < bufferLen[k] {
			max = bufferLen[k]
		}
	}
	return max / 2 + 1
}

/**
上面方法有些重复代码, 实现后的精简版
 */
func MaxLcpsLength1(s string) int {
	var buffer []string
	for _, v := range s {
		buffer = append(buffer, string(v))
	}
	newstring := strings.Join(buffer, "#")

	c := -1
	r := -1
	bufferLen := make([]int, len(newstring))
	max := -1
	for k, _ := range newstring {
		bufferLen[k] = 1
		i := 1
		if k < r {
			bufferLen[k] = bufferLen[2 * c - k]
			i = r - k
			if i < (bufferLen[2*c-k] / 2) {
				i = bufferLen[2*c-k] / 2
			}
		}
		for ; k - i >= 0 && k + i < len(newstring) && newstring[k - i] == newstring[k + i]; i ++ {
			bufferLen[k] = bufferLen[k] + 2
			if k + i > r {
				r = k + i
				c = k
			}
		}
		if max < bufferLen[k] {
			max = bufferLen[k]
		}
	}
	return max / 2 + 1
}

/**
返回最长回文子串
 */
func LongestPalindrome(s string) string {
	var buffer []byte
	for _, v := range s {
		buffer = append(buffer, '#', byte(v))
	}
	buffer = append(buffer, '#')
	newstring := string(buffer)
	//必须将ab扩展成 #a#b#的形式
	//1. 为什么扩展? 是因为这样能是字符串长度变成奇数, 方便原字符串为偶数的做比较
	//2. 为什么前后也有#, 因为对于abb来说, 如果变成a#b#b的话, 它会取#b#, b#b这两种歧义都是长度为3, 但原字符为1或2
	//如果变成#a#b#b#, 那就最长的是#b#b, 原字符为2

	c := -1
	r := -1
	bufferLen := make([]int, len(newstring))
	maxK := 0
	for k, _ := range newstring {
		//默认是进行暴力扩的, 所以长度为1
		bufferLen[k] = 1
		i := 1
		if k < r {
			//使用manacher进行的
			//比较对称点的回文半径哪个比较小就取哪个
			//涵盖了第2种的所有情况
			if bufferLen[2 * c - k] / 2 < r - k {
				bufferLen[k] = bufferLen[2 * c - k]
			}else {
				bufferLen[k] = (r - k) * 2 + 1
			}
			//半径从i开始扩展
			i =  bufferLen[k] / 2 + 1
		}
		//这里开始扩张
		for ; k-i >= 0 && k+i < len(newstring) && newstring[k-i] == newstring[k+i]; i++ {
			bufferLen[k] = bufferLen[k] + 2
			if k+i > r {
				r = k + i
				c = k
			}
		}
		if bufferLen[maxK] < bufferLen[k] {
			maxK = k
		}
	}

	//输出最长的字符串
	var out []byte
	for i := 0; i < bufferLen[maxK]; i++ {
		if string(newstring[maxK - bufferLen[maxK] / 2 + i]) == "#" {
			continue
		}
		out = append(out, newstring[maxK - bufferLen[maxK] / 2 + i])
	}
	return string(out)
}