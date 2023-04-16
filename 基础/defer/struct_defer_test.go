/**
  @author: wangyingjie
  @since: 2023/4/15
  @desc:
**/

package defer_test

import (
	"fmt"
	"testing"
)

type Person struct {
	age int
}

// TestStruct
//  @Description: 输出: 29 29 28
//  @param t
func TestStruct1(t *testing.T) {
	person := &Person{28}

	// 1.此时是将 28 当做 defer 函数的参数，会把 28 缓存在栈中，等到最后执行该 defer 语句的时候取出，即输出 28
	defer fmt.Println(person.age)

	// 2.defer 缓存的是结构体 Person{28} 的地址，最终 Person{28} 的 age 被重新赋值为 29，所以 defer 语句最后执行的时候，依靠缓存的地址取出的 age 便是 29，即输出 29
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3. 闭包引用，输出 29
	defer func() {
		fmt.Println(person.age)
	}()

	person.age = 29
}

// TestStruct2
//  @Description: 输出 29 28 28
//  @param t
func TestStruct2(t *testing.T) {
	person := &Person{28}

	// 1. 缓存28
	defer fmt.Println(person.age)

	// 2. defer 缓存的是结构体 Person{28} 的地址，这个地址指向的结构体没有被改变，最后 defer 语句后面的函数执行的时候取出仍是 28；
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3. 闭包引用，person 的值已经被改变，指向结构体 Person{29}，所以输出 29
	defer func() {
		fmt.Println(person.age)
	}()

	person = &Person{29}
}
