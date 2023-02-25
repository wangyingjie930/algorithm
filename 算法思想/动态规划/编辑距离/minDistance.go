/**
  @author: wangyingjie
  @since: 2023/2/25
  @desc: https://leetcode.cn/problems/edit-distance/
**/

package 编辑距离

import (
	"fmt"
	"math"
)

func minDistance(word1 string, word2 string) int {
	return helper(word1, word2, 0, 0)
}

func helper(source string, target string, sourceIndex, targetIndex int) int {
	if source == target {
		return 0
	}

	if len(source)-1 < sourceIndex || len(target)-1 < targetIndex {
		return int(math.Abs(float64(len(target) - len(source))))
	}

	if source[sourceIndex] == target[targetIndex] {
		return helper(source, target, sourceIndex+1, targetIndex+1)
	}

	//插入
	addString := source[0:sourceIndex] + string(target[targetIndex]) + source[sourceIndex:]

	//删除
	delString := source[0:sourceIndex] + source[sourceIndex+1:]

	//更新
	updateString := source[0:sourceIndex] + string(target[targetIndex]) + source[sourceIndex+1:]

	min := helper(addString, target, sourceIndex+1, targetIndex+1) + 1

	tmp := helper(delString, target, sourceIndex, targetIndex+1) + 1
	if tmp < min {
		min = tmp
	}

	tmp = helper(updateString, target, sourceIndex+1, targetIndex+1) + 1
	if tmp < min {
		min = tmp
	}
	fmt.Println(addString, delString, updateString, target, min)
	return min
}
