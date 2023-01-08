/**
  @author: wangyingjie
  @since: 2022/11/27
  @desc: consistentTest
**/

package consistentTest

import (
	"fmt"
	"testing"
)

func TestConsistent_Add(t *testing.T) {
	ips := []string{"192.168.1.1", "192.168.1.2", "192.168.1.3",
		"192.168.2.3", "192.168.2.4", "192.168.3.5", "192.168.3.6"}
	consist := NewConsistent()
	for _, v := range ips {
		consist.Add(v)
	}
	fmt.Printf("%v\n", consist)
	fmt.Println(consist.Get("hello"))
	fmt.Println(consist.Get("1qwe1"))
}
