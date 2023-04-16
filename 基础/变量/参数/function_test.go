/**
  @author: wangyingjie
  @since: 2023/4/16
  @desc:
**/

package function_test

import (
	"fmt"
	"testing"
)

// Test_FuncArr
//  @Description: 当传入已有的切片, 函数会在内部直接使用这个传入的切片，并不会创建一个的新的
//  @see https://studygolang.com/articles/11965
//  @param t
func Test_FuncArr(t *testing.T) {
	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0]) //等于18
}
func hello(num ...int) {
	num[0] = 18
}

// TestAppend
//  @Description: 函数可变参数配合append的探讨
//  @param t
func TestAppend(t *testing.T) {
	slice := []int{1, 2, -1, 0, 0}
	//change中的s和slice共用
	//当在change中进行append, 此时s超过长度, 已经开辟另外的内存存放了, 不会动到原始的slice
	change(slice...)
	t.Log(slice)

	//change中的s和slice共用
	//当在change中进行append, 没有超过长度, 在原始的slice中进行覆盖
	// https://www.topgoer.cn/docs/golang/chapter03-11
	change(slice[0:2]...)
	t.Log(slice)
}
func change(s ...int) {
	s[0] = 1000
	s = append(s, 3)
}

// funcMui
//  @Description: func funcMui(x,y int)(sum int, error) // 错误, 只要有一个返回值有命名，其他的也必须命名
//  @param x
//  @param y
//  @return sum
//  @return err
func funcMui(x, y int) (sum int, err error) {
	return x + y, nil
}

func TestAssert(t *testing.T) {
	x := 1
	justifyType(x)
}

// justifyType
//  @Description: 使用断言的x的类型需要为接口类型
//  @param x
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
