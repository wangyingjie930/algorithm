/**
  @author: wangyingjie
  @since: 2023/4/16
  @desc:
**/

package type_test

import (
	"fmt"
	"testing"
)

func Test_ArrayEqual(t *testing.T) {
	a := [2]int{5, 6}
	b := [2]int{5, 7}
	////如果b := [3]int{5, 7, 8}进行比较会发生报错,数组的长度也是数组类型的组成部分，所以 a 和 b 是不同的类型，是不能比较的
	if a == b {
		//数组也是只能比较是否相等, 并不能比较>或<
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}

func TestRangeChangeArray(t *testing.T) {
	var a = [5]int{1, 2, 3, 4, 5} //注意: 前提说的是数组
	var r [5]int

	for i, v := range a { // 这里参与的a是副本, 不是原本的a, 与a不是同一块内存
		//因此无论a被 如何修改，其副本a'依旧保持原值
		//可以使用for i, v := range &a 试试其他效果
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
	/**
	输出:
		r =  [1 2 3 4 5]
		a =  [1 12 13 4 5]
	*/
}
