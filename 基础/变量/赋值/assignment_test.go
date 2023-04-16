/**
  @author: wangyingjie
  @since: 2023/4/16
  @desc:
**/

package assignment_test

import (
	"fmt"
	"testing"
)

// TestSequence
//  @Description: 赋值顺序问题
//  @param t
func TestSequence(t *testing.T) {
	i := 1
	s := []string{"A", "B", "C"}
	/**
	1. 计算等号左边的索引表达式和取址表达式，接着计算等号右边的表达式；
	2. 赋值
	相当于  i, s[0] = 2, “Z”
	*/
	i, s[i-1] = 2, "Z"
	fmt.Printf("s: %v \n", s)
}

// TestDeclaration
//  @Description: ":="的声明问题
//  @param t
func TestDeclaration(t *testing.T) {
	var x *int
	//x, _ := foo() //error, 已经声明过了x
	//x, y = foo() //error, 没有声明过y

	x, _ = foo()  //对
	x, y := foo() //对, 多值赋值时，:= 左边的变量无论声明与否都可以
	fmt.Println(x, y)
}

type info struct {
	result int
}

func work() (int, error) {
	return 13, nil
}

// TestDeclarationStruct
//  @Description: 关于结构体的属性赋值问题
//  @param t
func TestDeclarationStruct(t *testing.T) {
	var data info

	//data.result, err := work() //error,不能使用短变量声明设置结构体字段值
	var err error
	data.result, err = work() //ok

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("info: %+v\n", data) //prints: info: {result:13}
}

// TestAssignNil
//  @Description: nil 只能赋值给指针、chan、func、interface、map 或 slice 类型的变量
//  @param t
func TestAssignNil(t *testing.T) {
	//var x = nil //error
	var x interface{} = nil
	//var x string = nil // error
	//var x error = nil //success,

	fmt.Println(x)
}
