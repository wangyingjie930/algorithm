/**
  @author: wangyingjie
  @since: 2023/3/11
  @desc: copy from chatgpt
  生产者消费者模型是一种常见的并发编程模型，它通常用于解决生产者和消费者之间的协作问题。
  在该模型中，生产者负责生产一些产品并将其放入共享的缓冲区中，而消费者则从缓冲区中取出产品并消费它们。
**/
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func producer(ch chan<- int) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Println("生产者生产了", i)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	defer wg.Done()
	for i := range ch {
		fmt.Println("消费者消费了", i)
	}
}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go producer(ch)
	go consumer(ch)
	wg.Wait()
}
