package main

import "fmt"

func BF(str, search string) int {
	j := 0
	for k, _ := range str {
		i := k
		for j < len(search) {
			if str[i] == search[j] {
				j ++
				i ++
			}else {
				j = 0
				break
			}
		}
		if j == len(search) {
			return k
		}
	}
	return -1
}

func KMP(str string, search string) int {
	i := 0
	j := 0
	nextArr := getNextArr(search)
	for i < len(str) && j < len(search) {
		if str[i] == search[j] {
			i ++
			j ++
		}else if nextArr[j] == -1 {
			i ++
		}else {
			j = nextArr[j]
		}
	}
	if j == len(search) {
		return i - j
	}
	return -1
}

func getNextArr(str string) []int {
	nextArr := make([]int, len(str))
	nextArr[0] = -1
	nextArr[1] = 0
	i := 2
	j := 0
	for i < len(str) {
		if str[j] == str[i - 1] {
			//将当前的上一个字符与上一次的比较字符(同时也是最长长度)做比较
			j ++
			nextArr[i] =j
			i ++
		}else if j > 0 {
			//类似KMP, 移动指针
			j = nextArr[j]
		}else {
			//找不到, 则为0
			nextArr[i] = 0
			i ++
		}
	}
	return nextArr
}

func main() {
	//fmt.Print(getNextArr("sasasa"))
	fmt.Println(KMP("adsasadsasasadd", "sasasa"))
	fmt.Println(BF("adsasadsasasadd", "sasasa"))
}
