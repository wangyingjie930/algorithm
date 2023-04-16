package main

import "fmt"

type Animal interface { //接口
	say()
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func (_self Dog) say() {
	fmt.Println("I am", _self.name)
}

func (_self *Cat) say() {
	fmt.Println("I am", _self.name)
}

// main
//  @Description:
// 1、若以实体类型（T）实现接口，不管是T类型的值，还是T类型的指针，都实现了该接口。
// 2、若以指针类型（*T）实现接口，只有*T实现了该接口。
func main() {
	var d1 Animal = Dog{name: "wangCai1"}
	d1.say() //因为Dog 的value类型实现了Animal接口

	var d2 Animal = &Dog{name: "wangCai2"}
	d2.say() //因为dDog 的指针类型实现了Animal接口

	//var c1 Animal= Cat{name:"miaoMiao1"}
	//c1.say() //因为Cat 的value类型没有实现了Animal接口，所以报错

	var c2 Animal = &Cat{name: "miaoMiao2"}
	c2.say() //因为Cat 的指针类型实现了Animal接口
}
