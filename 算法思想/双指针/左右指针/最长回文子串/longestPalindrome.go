/**
  @author: wangyingjie
  @since: 2023/2/27
  @desc: https://leetcode.cn/problems/longest-palindromic-substring/description/
**/

package 最长回文子串

func longestPalindrome(s string) string {
	longest := ""
	for i := 0; i < len(s); i++ {
		s1 := Palindrome(s, i, i)
		s2 := Palindrome(s, i, i+1)
		if len(s1) > len(longest) {
			longest = s1
		}
		if len(s2) > len(longest) {
			longest = s2
		}
	}
	return longest
}

// Palindrome
//  @Description: 由中心点向外部扩散
//  @param s
//  @param i
//  @param j
//  @return string
func Palindrome(s string, i, j int) string {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return s[i+1 : j]
}
