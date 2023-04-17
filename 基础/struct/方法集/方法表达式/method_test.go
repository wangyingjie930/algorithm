/**
  @author: wangyingjie
  @since: 2023/4/17
  @desc: 通过类型引用的方法表达式会被还原成普通函数样式
**/

package method_test

import (
	"fmt"
	"testing"
)

type N int

func (n N) test() {
	fmt.Println(n)
}

func TestMethod(t *testing.T) {
	var n N = 10
	fmt.Println(n)

	n++
	f1 := N.test
	f1(n)

	n++
	f2 := (*N).test
	f2(&n)
}

func TestMethod1(t *testing.T) {
	var n N = 10
	fmt.Println(n)

	n++
	f1 := N.test
	f1(n)

	n++
	f2 := (*N).test
	f2(&n)
}
