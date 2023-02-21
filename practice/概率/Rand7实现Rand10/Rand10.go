/**
  @author: wangyingjie
  @since: 2023/2/21
  @desc: https://leetcode.cn/problems/implement-rand10-using-rand7/solutions/
**/

package Rand7实现Rand10

import (
	"math/rand"
	"time"
)

func rand7() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()%7 + 1
}

func rand10() int {
	//每调用一次rand7相当于抛了一次硬币
	//调用两次的情况有11, 12,...., 17, 21....27, 31, ..., 77
	//等价于7进制了 00, 01,...., 06, 10, ..., 66

	//所以当要实现rand10, 覆盖的范围是[1, 10], 则需选择2位的7进制数才能满足大于10的范围
	//所以需要调用2次rand7, 而且要满足7进制的写法
	for true {
		num := (rand7()-1)*7 + (rand7() - 1)
		if num <= 10 && num >= 1 {
			return num
		}
	}
	return -1
}
