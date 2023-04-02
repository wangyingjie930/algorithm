/**
  @author: wangyingjie
  @since: 2023/4/2
  @desc: https://leetcode.cn/problems/word-break/description/
**/

package 单词拆分

// wordBreak
//  @Description: 递归+备忘录实现
//  @param s
//  @param wordDict
//  @return bool
func wordBreak(s string, wordDict []string) bool {
	memo := make(map[int]bool, len(s))
	return dfs(s, 0, wordDict, memo)
}

func dfs(s string, i int, wordDict []string, memo map[int]bool) bool {
	if i == len(s) {
		return true
	}

	if _, exist := memo[i]; exist {
		return memo[i]
	}

	for j := 0; j < len(wordDict); j++ {
		if len(wordDict[j]) <= len(s)-i && wordDict[j] == s[i:i+len(wordDict[j])] {
			if dfs(s, i+len(wordDict[j]), wordDict, memo) {
				memo[j] = true
				return true
			}
		}
	}

	memo[i] = false
	return false
}

func wordBreakDp(s string, wordDict []string) bool {
	stringMapper := make(map[string]bool, len(wordDict))
	for i := 0; i < len(wordDict); i++ {
		stringMapper[wordDict[i]] = true
	}

	//dp[i]: 表示长度为i的字符串是否匹配
	//dp[i] = dp[j] && 字符串[j, i-1]的位置是否有字符串匹配 (这里的dp[j]实际的含义就是字符串[0, j-1]是否有字符串匹配)
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			dp[i] = dp[i] || (dp[j] && stringMapper[s[j:i]])
		}
	}
	return dp[len(s)]
}
