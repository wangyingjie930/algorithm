package BM

//坏字符规则-初始化散列表
func generateForBadHash(search string) map[int]int {
	ret := make(map[int]int)
	for k, c := range search {
		ret[int(c)] = k
	}
	return ret
}

//坏字符匹配
func badCharRule(subject string, search string) int {
	bc := generateForBadHash(search)
	//subject和search以开头对齐
	i := 0
	//最起码也是尾部对齐, 所以i不超过字符串-匹配串的长度
	for ; i <= len(subject)-len(search); {
		j := len(search) - 1
		for ;j >= 0; {
			if subject[i + j] == search[j] {
				//遇到相同的j往下继续匹配
				j --
			}else {
				//遇到不相同的使用散列表将不相同的字符与匹配串的字符相对齐
				x := -1
				if v, ok := bc[int(subject[i + j])]; ok {
					x = v
				}
				i = i + j - x
				break
			}
		}
		//比较完毕, 返回i
		if j < 0 {
			return i
		}
	}
	return -1
}

//好字符规则-根据匹配串生成suffix, prefix数组
func generateGoodRule(search string) ([]int, []bool){
	suffix := make([]int, len(search))
	prefix := make([]bool, len(search))
	for i := 0; i < len(search); i++ {
		suffix[i] = -1
		prefix[i] = false
	}

	for i := 0; i < len(search) - 1; i ++ {
		j := i
		k := 0
		for ;j >= 0; {
			//将最后的子串与中间的子串进行对比
			//相等, 则suffix填入相同子串的长度及对应的索引
			if search[j] == search[len(search) - 1 - k] {
				k ++
				suffix[k] = j
				j --
			}else {
				break
			}
		}
		if j == -1 {
			prefix[k] = true
		}
	}
	return suffix, prefix
}

