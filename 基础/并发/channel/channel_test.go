/**
  @author: wangyingjie
  @since: 2023/4/16
  @desc:
**/

package channel

import (
	"fmt"
	"testing"
	"time"
)

func Stop(stop <-chan bool) {
	//close(stop) //error: 有方向的 channel 不可以被关闭
}

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

// TestInOutChan
//  @Description: 限制通道在函数中只能发送或只能接收
//  @param t
func TestInOutChan(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}

func TestChannelAssign(t *testing.T) {
	var ch chan int
	//ch = make(chan int)
	go func() {
		ch <- 1 //使用channel时不make的话,读写都会阻塞
		fmt.Println("tets")
	}()

	for {
		//用于不会跳出循环
		select {
		case v, ok := <-ch:
			println(v, ok)
			return
		default:
			println("default")
		}
		time.Sleep(1 * time.Second)
	}
}

var o = fmt.Print

// TestChannel
//  @Description: 输出 3 2 1
//  @param t
func TestChannel(t *testing.T) {
	c := make(chan int, 1)
	for range [3]struct{}{} {
		select {
		default:
			o(1)
		case <-c: //2. 可读
			o(2)
			c = nil //3. 置为nil, 后续不可读写
		case c <- 1: //1.可写
			o(3)
		}
	}
}
