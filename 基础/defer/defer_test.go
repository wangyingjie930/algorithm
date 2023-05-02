/**
  @author: wangyingjie
  @since: 2023/4/15
  @desc: Go 语言中所有的函数调用都是传值的，虽然 defer 是关键字，但是也继承了这个特性
**/

package defer_test

import (
	"fmt"
	"testing"
)

// deferCall
//  @Description:
// defer 的执行顺序是后进先出。当出现 panic 语句的时候，会先按照 defer 的后进先出的顺序执行，最后才会执行panic
func deferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}

func Test_defer(t *testing.T) {
	deferCall()
}

func Test_sequence(t *testing.T) {
	i := 5
	defer hello(i) //在执行 defer 语句的时候会保存一份副本，在实际调用 hello() 函数时用，所以是 5.
	i = i + 10
}
func hello(i int) {
	fmt.Println(i)
}

// Test_return
//  @Description: defer的调用顺序为return之后, 返回值返回之前
//  @param t
func Test_return(t *testing.T) {
	/**
	输出:
		test defer, i =  1
		test return: 1
	*/
	ret := returnTestWithoutInit()
	fmt.Println("test return:", ret)

	/**
	输出:
		test defer, i =  1
		test return: 0
	*/
	ret = returnTest()
	fmt.Println("test return:", ret)
}

// returnTestWithoutInit
//  @Description:
// 1. 先进行return
// 2. 调用了defer, 此时进行i++, 由于函数定义时的返回值是(i int), 此时更改i, 就修改了返回值, 所以返回1
//  @return i
func returnTestWithoutInit() (i int) {
	defer func() {
		i++
		fmt.Println("test defer, i = ", i)
	}()

	return i
}

// returnTest
//  @Description:
// 1. 先进行return
// 2. 调用了defer, 虽然在函数中修改了, 但是并没有返回给调用方
//  @return int
func returnTest() int {
	var i int
	defer func() {
		i++
		fmt.Println("test defer, i = ", i)
	}()

	return i
}

func Test_Return1(t *testing.T) {
	t.Log(f1()) //1
	t.Log(f2()) //10
	t.Log(f3()) //1
}

// f1
//  @Description:
// 1. 当执行到return后, 相当于r = 0进行语句赋值
// 2. 之后调用defer函数, 操作的是r这个变量, 所以r++, 返回的是1
//  @return r
func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

// f2
//  @Description:
// 1. 当执行到return后, 相当于r = t进行语句赋值
// 2. 之后调用defer函数, 只是在操作t, 并没有修改r的值
//  @return r
func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

// f3
//  @Description:
// 1. 当执行到return后, 相当于r = 1进行语句赋值
// 2. 之后调用defer函数, 传递参数r, 修改的是r的副本, 并没有修改r的值
//  @return r
func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

// TestPreReturn
//  @Description: 输出 2 1, 没有被注册的defer不会执行
//  @param t
func TestPreReturn(t *testing.T) {
	defer func() {
		fmt.Println("1")
	}()
	if true {
		fmt.Println("2")
		return
	}
	defer func() {
		fmt.Println("3")
	}()
}

// TestRecover
//  @Description:
// 1. 执行到return时, 相当于 r = n + 1 = 4
// 2. 调用了f()函数, (因为没有定义会报错)
// 3. 调用了第一个defer, r = r + n = 7, 并且recover了f()函数
//  @param t
func TestRecover(t *testing.T) {
	ret := func(n int) (r int) {
		defer func() {
			r += n
			recover()
		}()

		var f func()

		defer f()
		f = func() {
			r += 2
		}
		return n + 1
	}(3)

	t.Log(ret)
}

func TestTest(t *testing.T) {
	fmt.Println(test1())
	fmt.Println(test2())
	fmt.Println(test3())
	fmt.Println(test4())

	/**
	0, 0, 3, 3, 0, 4, 0, 5
	*/

	return
}

func test1() (v int) {
	defer fmt.Println(v)
	return v
}

func test2() (v int) {
	//defer关键字时也使用值传递，但是因为拷贝的是函数指针, 所以能找到最终的v
	defer func() {
		fmt.Println(v)
	}()
	return 3
}

func test3() (v int) {
	//此时的v已经初始化, v=0
	//调用 defer 关键字会立刻拷贝函数中引用的外部参数!!!
	defer fmt.Println(v) //输出为0
	v = 3
	return 4
}

func test4() (v int) {
	defer func(n int) {
		fmt.Println(n)
	}(v) //类似test3, 此时的v已经初始化, v=0, 保存在里面了, 后续输出为0
	return 5
}
