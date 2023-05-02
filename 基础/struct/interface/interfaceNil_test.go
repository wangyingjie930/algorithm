/**
  @author: wangyingjie
  @since: 2023/4/16
  @desc:
**/

package interface_test

import (
	"fmt"
	"testing"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func TestNil(t *testing.T) {
	var s *Student
	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not nil")
	}
	//记住一点，当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil
	//给变量 p 赋值之后，p 的动态值是 nil，但是动态类型却是 *Student，是一个 nil 指针
	var p People = s
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
}

func Test_InterfaceNil(t *testing.T) {
	var i interface{} //当且仅当接口的动态值和动态类型都为 nil 时，接口类型值才为 nil
	if i == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println("not nil")
}

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

// TestNonEmptyInterface
//  @Description: 调用Foo时相当于把*int的动态类型赋给了Foo中的x, 所以Foo中的x并不等于nil
//  (当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil)
//  @param t
func TestNonEmptyInterface(t *testing.T) {
	var x *int = nil
	Foo(x)
}

// TestNil1
//  @Description:https://www.topgoer.cn/docs/gomianshiti/mian43
//  @param t
func TestNil1(t *testing.T) {
	x := interface{}(nil)
	y := (*int)(nil) //相当于var y *int, 因为指针类型给零值是nil(对于一个接口的零值就是它的类型和值的部分都是nil)
	a := y == x
	b := y == nil
	//类型断言语法：i.(Type)，其中 i 是接口，Type 是类型或接口。
	//但是，如果动态类型不存在，则断言总是失败
	_, c := x.(interface{})
	println(a, b, c, x == nil) //false true false true

	var d interface{} = y
	var e interface{} = x
	var f *int
	println(d == nil, e == nil, y == f) //false true true
}
