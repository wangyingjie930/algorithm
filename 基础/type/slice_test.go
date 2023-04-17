/**
  @author: wangyingjie
  @since: 2023/4/15
  @desc:
**/

package type_test

import (
	"fmt"
	"testing"
	"time"
)

func TestZero(t *testing.T) {
	/**
	关于零值的
	*/
	var s1 []int     //零值, 引用类型为nil
	var s2 = []int{} //这是空切片
	if s1 == nil {   //相等
		t.Log("s1 is nil")
	}
	if s2 == nil { //空切片不等于nil
		t.Log("s2 is nil")
	}
}

func TestCap(t *testing.T) {
	/**
	关于长度及容量
	*/
	s := [3]int{1, 2, 3}
	a := s[:0]         // cap(a) = 3[s的容量] - 0[切片起始位置] = 3
	b := s[:2]         //cap(b) = 3[s的容量] - 0[切片起始位置] = 3
	c := s[1:2:cap(s)] // cap(c) = 3[s的容量] - 1[切片起始位置] = 2
	t.Log("cap:", cap(a), cap(b), cap(c))
	t.Log("slice:", a, b, c, s[:3])

	//假设基础切片是 baseSlice，使用操作符 [low,high]，有如下规则：0 <= low <= high <= cap(baseSlice)
	//s[:4] //error, 超过原本数组的容量
}

func TestSubSlice(t *testing.T) {
	x := make([]int, 2, 10)
	_ = x[6:10]
	_ = x[6:] //截取符号 [i:j]，如果j省略，默认是原切片或者数组的长度, 相当于x[6:2]
	_ = x[2:]
}

func TestAppend(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	s2[1] = 4
	fmt.Println(s1) //[1 2 4], 因为是引用类型嘛
	s2 = append(s2, 5, 6, 7)
	fmt.Println(s1) //[1 2 4], append s2数组时会使s2生成新的数组
}

func TestGoRange(t *testing.T) {
	var m = [...]int{1, 2, 3}

	for i, v := range m {
		//变量 i、v 在每次循环体中都会被重用，而不是重新声明
		//循环结束遍历一般比goroutine更快结束
		//当循环结束时, goroutine拿到的i/v都是同一个位置的值, 里面的值就是循环之后最后存放的值
		go func() {
			fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Second * 3)
	/**
	输出:
		2, 3
		2, 3
		2, 3
	*/
}

func TestGoRange1(t *testing.T) {
	var m = [...]int{1, 2, 3}

	for i, v := range m {
		//变量 i、v 在每次循环体中都会被重用，而不是重新声明
		go func() {
			fmt.Println(i, v)
		}()
		//这里相较于之前加了一个时间
		//那么goroutine就可以在循环为结束之前获得v位置的值并输出里面存放的值
		time.Sleep(time.Second * 1)
	}
	time.Sleep(time.Second * 3)
	/**
	输出:
		0 1
		1 2
		2 3
	*/
}

func TestGoRange2(t *testing.T) {
	var m = [...]int{1, 2, 3}

	for i, v := range m {
		//变量 i、v 在每次循环体中都会被重用，而不是重新声明
		//循环结束遍历一般比goroutine更快结束
		//当循环结束时, goroutine拿到的i/v都是不同位置的值, 因为有:=进行重新分配
		i := i
		v := v
		go func() {
			fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Second * 3)
	/**
	输出:[顺序可能不一致]
		2 3
		1 2
		0 1
	*/
}

// TestRangeChangeSlice
//  @Description:https://tonybai.com/2015/09/17/7-things-you-may-not-pay-attation-to-in-go/
// 记住一个核心, range遍历就是复制原数据的副本之后再进行遍历
//  @param t
func TestRangeChangeSlice(t *testing.T) {
	var a = []int{1, 2, 3, 4, 5} //注意: 前提说的是切片
	var r [5]int

	for i, v := range a {
		//它实际上复制的是一个slice，也就是那个struct
		//副本struct中的*T依旧指向原slice对应的array，为此对slice的修改都反映到 underlying array a上去了
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
	/**
	输出:
		r =  [1 12 13 4 5]
		a =  [1 12 13 4 5]
	*/
}

// TestRange
//  @Description: 可以结束循环
//  @param t
func TestRangeSliceLen(t *testing.T) {
	v := []int{1, 2, 3}
	for i := range v {
		//复制了v这个切片的副本, 然后这个切片的struct中的len字段并没有发生改变
		//因此对于原切片v的长度的增大不会影响到副本v的长度
		v = append(v, i)
	}
}

func TestPointer(t *testing.T) {
	arrayA := [5]int{100, 200}
	arrayB := arrayA[:]
	i := 1

	//表示存放i位置的地址
	//%p针对的类型是指针, fmt.Printf("i : %p\n", i)是错误的
	//使用&i得到指针类型, 内容保存的是i的地址, %p以16进制的形式输出保存地址的内容
	fmt.Printf("i : %p\n", &i)

	//表示存放数组的地址, 同理
	fmt.Printf("arrayA : %p\n", &arrayA)

	//arrayB---->数组
	//&arrayB表示存放切片的地址, %p输出的是切片的位置内容
	//arrayB本身就是引用, 本身保存的就有数组的位置信息, %p输出数组的位置
	fmt.Printf("arrayB : %p , %p\n", &arrayB, arrayB)
}

func TestCopy(t *testing.T) {
	array := []int{10, 20, 30, 40}
	slice := make([]int, 6)
	n := copy(slice, array)
	fmt.Println(n, slice)
}

// TestIndex
//  @Description: 没有指定索引的元素会在前一个索引基础之上加一
// 相当于 []int{2: 2, 3: 3, 0: 1}
//  @param t
func TestIndex(t *testing.T) {
	var x = []int{2: 2, 3, 0: 1}
	fmt.Println(x) //输出: [1 0 2 3]
}
