/**
  @author: wangyingjie
  @since: 2023/3/23
  @desc: https://leetcode.cn/problems/reverse-words-in-a-string/
**/

package 反转字符串中的单词

func reverseWords(s string) string {
	arr := trimString(s)              //清除多余空格
	reverseString(arr, 0, len(arr)-1) //先整体翻转一遍

	for i := 0; i < len(arr); {
		j := i
		for j < len(arr) && arr[j] != ' ' {
			j++
		}
		reverseString(arr, i, j-1) //根据空格切分, 进行反转操作
		i = j + 1
	}
	return string(arr)
}

func trimString(s string) []byte {
	var ret []byte

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			for i < len(s) && s[i] == ' ' {
				i++
			}
			if i < len(s) && len(ret) > 0 {
				ret = append(ret, ' ')
			}
		}
		if i < len(s) {
			ret = append(ret, s[i])
		}
	}
	return ret
}

func reverseString(s []byte, start, end int) {
	for start <= end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
