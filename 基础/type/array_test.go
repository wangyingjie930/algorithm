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

func TestArrayNil(t *testing.T) {
	var k = 9
	for k = range []int{} { //空切片不进入循环
		t.Log("enter1")
	}
	fmt.Println(k)

	for k = 0; k < 3; k++ {
	}
	fmt.Println(k)

	te := (*[3]int)(nil)

	for k = range te { //这里能进入循环, 就奇了怪了
		t.Log("enter2")
	}
	fmt.Println(k)
	//输出: 9 3 2 enter2 enter2 enter2
}
