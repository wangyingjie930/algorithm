package main

import "fmt"

// 旧接口定义
type Greeter interface {
	Greet() string
}

// 新接口定义，嵌套原有的接口类型
type Talker interface {
	Greeter // 嵌套 Greeter 接口类型
	Talk() string
}

// User 结构体类型实现了 Greet 方法和 Talk 方法
type User struct {
	Name string
}

func (u *User) Greet() string {
	return fmt.Sprintf("Hello, my name is %s", u.Name)
}

func (u *User) Talk() string {
	return fmt.Sprintf("%s, nice to meet you!", u.Greet())
}

func main() {
	// 创建 User 类型实例
	u := &User{"Alice"}

	// 将 User 类型的指针赋值给 Talker 接口类型
	var talker Talker = u

	// 输出结果为 "Hello, my name is Alice, nice to meet you!"
	fmt.Println(talker.Talk())
}
