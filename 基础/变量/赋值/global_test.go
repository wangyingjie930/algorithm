/**
  @author: wangyingjie
  @since: 2023/4/16
  @desc: 全局变量重新赋值问题
**/

package assignment_test

import (
	"fmt"
	"testing"
)

var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	fmt.Println(*p)
}

func TestRecover(t *testing.T) {
	//问题出在操作符:=，对于使用:=定义的变量，如果新变量与同名已定义的变量不在同一个作用域中，那么 Go 会新定义这个变量
	p, err := foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	//调用bar的时候会报错
	//p 是新定义的变量，会遮住全局变量 p，导致执行到bar()时程序，全局变量 p 依然还是 nil，程序随即 Crash
	bar()
	fmt.Println(*p)
}

var str string //var str = ""
//str := "" //error, 全局变量不能使用:=进行
