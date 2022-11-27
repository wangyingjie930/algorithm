package main

import (
	"fmt"
	consistentTest "imooc-product/backend/一致性哈希算法"
)

func main() {
	ips := []string{"192.168.1.1", "192.168.1.2", "192.168.1.3",
		"192.168.2.3", "192.168.2.4", "192.168.3.5", "192.168.3.6"}
	consist := consistentTest.NewConsistent()
	for _, v := range ips {
		consist.Add(v)
	}
	fmt.Printf("%v\n", consist)
	fmt.Println(consist.Get("hello"))
	fmt.Println(consist.Get("1qwe1"))
}
