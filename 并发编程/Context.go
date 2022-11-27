package 并发编程

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type myKey int

func addNum(numP *int32, id int, deferFunc func()) {
	defer func() {
		deferFunc()
	}()
	for i := 0; ; i++ {
		currNum := atomic.LoadInt32(numP)
		newNum := currNum + 1
		time.Sleep(time.Millisecond * 200)
		//旧值指针和旧值传递时的值是否相同, 相同表示内存没有被其他线程改动过
		//此时就将newNum赋值给numP
		if atomic.CompareAndSwapInt32(numP, currNum, newNum) {
			fmt.Printf("The number: %d [%d-%d]\n", newNum, id, i)
			break
		} else {
			//fmt.Printf("The CAS operation failed. [%d-%d]\n", id, i)
		}
	}
}

func coordinateWithWaitGroup() {
	total := 12
	stride := 3
	var num int32
	fmt.Printf("The number: %d [with sync.WaitGroup]\n", num)
	var wg sync.WaitGroup
	for i := 1; i <= total; i = i + stride {
		wg.Add(stride)
		for j := 0; j < stride; j++ {
			go addNum(&num, i+j, wg.Done)
		}
		wg.Wait()
	}
	fmt.Println("End.")
}

func coordinateWithContext() {
	total := 12
	var num int32
	fmt.Printf("The number: %d [with context.Context]\n", num)
	//context.Background()提供根content
	//WithCancel“撤销”这个操作是Context值能够协调多个 goroutine 的关键所在
	cxt, cancelFunc := context.WithCancel(context.Background())
	for i := 1; i <= total; i++ {
		go addNum(&num, i, func() {
			if atomic.LoadInt32(&num) == int32(total) {
				//手动撤销
				cancelFunc()
			}
		})
	}
	//这里也要使用
	<-cxt.Done()
	fmt.Println("End.")
}
