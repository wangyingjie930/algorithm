package KMP

func getIndexOf(subject, search string) int {

	i := 0
	j := 0
	//获取next数组
	nextArr := getNextArr(search)

	for i < len(subject) && j < len(search) {
		if subject[i] == search[j] {
			i ++
			j ++
		}else if j == 0 {
			i ++
		}else {
			//记住核心结论: next[j]就是下次要指向的search串的位置
			//同时这里就是起到加速的地方
			j = nextArr[j]
		}
	}

	if j >= len(search) {
		//表示已经比对成功了, 才能够走到search串结束
		return i - len(search)
	}
	//表示search都没走完, subject就走完了, 表示找不到
	return -1
}

func getNextArr(search string) []int {
	if len(search) == 1 {
		return []int{-1}
	}
	nextArr := make([]int, len(search), len(search))
	//这是固定的
	nextArr[0] = -1
	nextArr[1] = 0
	//从2开始
	i := 2
	//当前的cn是累计的next值, 取的是next[i - 1] => next[1] = 0
	//一方面表示的是相同前缀字符串的长度, 一方面又是应该回跳的位置
	cn := 0
	for ; i < len(search); {
		if search[cn] == search[i-1] {
			//实际就是nextArr[i] = next[i - 1] + 1
			nextArr[i] = cn + 1
			i ++
			cn ++
		} else if cn > 0 {
			//同KMP的结论一样, 下一个指向是放在j = nextArr[j]
			cn = nextArr[cn]
		}else {
			//当以上条件都不满足时, cn已经到了0, 没法继续回退了, nextArr当前值为0
			nextArr[i] = 0
			//往下走
			i ++
		}
	}
	return nextArr
}
