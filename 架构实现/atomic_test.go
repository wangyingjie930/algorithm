package 架构实现

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func Test_Atomic(t *testing.T) {

	// 示例1。
	var box atomic.Value
	fmt.Println("Copy box to box2.")
	box2 := box // 原子值在真正使用前可以被复制。
	v1 := [...]int{1, 2, 3}
	fmt.Printf("Store %v to box.\n", v1)
	box.Store(v1)
	fmt.Printf("The value load from box is %v.\n", box.Load())
	fmt.Printf("The value load from box2 is %v.\n", box2.Load())
	fmt.Println()

	// 示例3。
	fmt.Println("Copy box to box3.")
	box3 := box // 原子值在真正使用后不应该被复制！
	fmt.Printf("The value load from box3 is %v.\n", box3.Load())
	v3 := 123
	fmt.Printf("Store %d to box3.\n", v3)
	//box3.Store(v3) // 这里会引发一个panic，报告存储值的类型不一致。
	_ = box3
	fmt.Println()

	// 示例6。
	var box6 atomic.Value
	v6 := []int{1, 2, 3}
	fmt.Printf("Store %v to box6.\n", v6)
	box6.Store(v6)
	v6[1] = 4 // 注意，此处的操作不是并发安全的！
	fmt.Printf("The value load from box6 is %v.\n", box6.Load())
	// 正确的做法如下。
	v6 = []int{1, 2, 3}
	store := func(v []int) {
		replica := make([]int, len(v))
		copy(replica, v)
		box6.Store(replica)
	}
	fmt.Printf("Store %v to box6.\n", v6)
	store(v6)
	v6[2] = 5 // 此处的操作是安全的。
	fmt.Printf("The value load from box6 is %v.\n", box6.Load())
}
