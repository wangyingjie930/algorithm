/**
  @author: wangyingjie
  @since: 2023/4/17
  @desc:
**/

package mutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type data struct {
	*sync.Mutex // *Mutex, 这里的指针引用很重要!!, 涉及到后面拿到是不是同位置的锁
}

func (d data) test(s string) { // 值方法
	d.Lock() //2. 这里调用了lock方法, d结构体传值, 所以*sync.Mutex的地址信息也拿到了, 用的是同块内存的锁
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}

func TestMutex(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	d := data{new(sync.Mutex)} // 初始化

	go func() {
		defer wg.Done()
		d.test("read") //1. 在这里相当于结构体传值
	}()

	go func() {
		defer wg.Done()
		d.test("write")
	}()

	wg.Wait()
}
