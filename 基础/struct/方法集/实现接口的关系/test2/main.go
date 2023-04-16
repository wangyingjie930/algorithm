package main

import "fmt"

type Animal interface {
	say()
	doSome()
}

type Dog struct {
	name string
}

func (_self Dog) say() {
	//表示实现了*Dog还有Dog的方法
	fmt.Println("I am", _self.name)
}
func (_self *Dog) doSome() {
	//只实现了*Dog的方法
	fmt.Println("I will do something")
}

//所以是*Dog实现了Animal的接口

func main() {
	// 报错，因为Dog的value类型实现了Animal接口的say方法没有实现doSome方法
	//var d1 Animal= Dog{name:"wangCai1"}
	//d1.say()

	//因为*Dog实现了Animal接口集的所有方法
	var d2 Animal = &Dog{name: "wangCai2"}
	d2.say()
}
