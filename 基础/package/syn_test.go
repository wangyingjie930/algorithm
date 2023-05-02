package go_package

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSynMap(t *testing.T) {
	var sMap sync.Map

	sMap.Store("key1", "value1")
	sMap.Store(2, 2)

	sMap.Range(func(key, value interface{}) bool {
		t.Log(key, value)
		return true
	})

	t.Log(sMap.Load(3))
}

func TestSynCond(t *testing.T) {
	type Account struct {
		balance int
		lock    sync.Mutex
		cond    *sync.Cond
	}

	account := &Account{}
	account.cond = sync.NewCond(&account.lock) //cond和lock成对出现, cond里面也会用到lock

	// 存款方法
	deposit := func(amount int) {
		account.lock.Lock()
		defer account.lock.Unlock()

		account.balance += amount
		fmt.Printf("存入 %d 元，当前余额为 %d 元\n", amount, account.balance)
		account.cond.Broadcast()
	}

	// 取款方法
	withdraw := func(amount int) bool {
		account.lock.Lock()
		defer account.lock.Unlock()

		for account.balance < amount {
			account.cond.Wait()
		}
		account.balance -= amount
		fmt.Printf("取出 %d 元，当前余额为 %d 元\n", amount, account.balance)
		return true
	}

	// 开启2个 goroutine 进行存、取款操作
	for i := 0; i < 2; i++ {
		go func() {
			deposit(100)
		}()

		go func() {
			withdraw(150)
		}()
	}
	time.Sleep(3 * time.Second)
}

func TestSyncPool(t *testing.T) {
	// 定义一个结构体，用于将需要重用的对象存放到临时对象池中
	type MyObject struct {
		id int // 对象标识符
	}

	var pool = sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new object")
			return &MyObject{}
		},
	}

	// 从池中获取对象
	obj1 := pool.Get().(*MyObject)
	// 设置对象属性及操作
	obj1.id = 1
	fmt.Printf("Object 1: %+v\n", obj1)
	// 将对象放回池中，供后续重用
	pool.Put(obj1)

	// 从池中获取对象
	obj2 := pool.Get().(*MyObject)
	// 由于上一次设置了 id 属性，因此这次重用的对象依旧有相同的 id
	fmt.Printf("Object 2: %+v\n", obj2)
	// 将对象放回池中，供后续重用
	pool.Put(obj2)
}
