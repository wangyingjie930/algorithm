package LRU缓存淘汰

import (
	"fmt"
	"testing"
)

func TestLRUCache_Put(t *testing.T) {
	//[[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]
	lru := Constructor(3)
	lru.Put(1, 1)
	fmt.Println(lru)
	lru.Put(2, 2)
	fmt.Println(lru)
	lru.Put(3, 3)
	fmt.Println(lru)
	lru.Put(4, 4)
	fmt.Println(lru)

	fmt.Println(lru.Get(4))
	fmt.Println(lru)
	fmt.Println(lru.Get(3))
	fmt.Println(lru)
	fmt.Println(lru.Get(2))
	fmt.Println(lru)
	fmt.Println(lru.Get(1))
	fmt.Println(lru)
}
