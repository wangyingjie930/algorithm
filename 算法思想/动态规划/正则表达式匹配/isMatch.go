/**
  @author: wangyingjie
  @since: 2023/5/18
  @desc: https://leetcode.cn/problems/regular-expression-matching/
**/

package isMatch

func isMatch(s string, p string) bool {
	return helper(s, 0, p, 0)
}

func helper(s string, i int, p string, j int) bool {
	if i == len(s)-1 && j == len(p)-1 {
		return true
	} else if i > len(s)-1 || j > len(p)-1 {
		return false
	}

	if s[i] == p[j] || p[j] == '.' {
		//当前对比的字符相同
		if j+1 < len(p) && p[j+1] == '*' {
			return helper(s, i, p, j+2) || //不进行匹配*
				helper(s, i+1, p, j) //尝试进行匹配*
		} else {
			return helper(s, i+1, p, j+1)
		}
	} else {
		//当前对比的字符不同
		if j+1 < len(p) && p[j+1] == '*' {
			return helper(s, i, p, j+2)
		} else {
			return false
		}
	}
}
