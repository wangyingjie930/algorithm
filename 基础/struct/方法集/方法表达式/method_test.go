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
	n = n + 1
	fmt.Println(n)
}

func TestMethod(t *testing.T) {
	var n N = 10
	fmt.Println(n)

	n++
	f2 := (*N).test
	f2(&n)

	n++
	f1 := N.test
	f1(n)
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

type A int

func (a A) Value() int {
	fmt.Printf("%+v\n", &a)
	return int(a)
}

func (a A) SetValue(n int) {
	a = A(n)
}

func (a *A) Set(n int) {
	*a = A(n)
}

type B struct {
	A
	b int
}

type C struct {
	*A
	b int
}

func TestA(t *testing.T) {
	var a A = 1
	var b B = B{a, 2}
	var b1 *B = &b
	var c C = C{&a, 2}
	var c1 *C = &C{&a, 2}

	(*A).Set(&b.A, 2)
	(*A).Set(&b1.A, 2)
	(*A).Set(c.A, 2)
	(*A).Set(c1.A, 2)
	fmt.Println(b.Value(), b1.Value(), c.Value(), c1.Value())
}
