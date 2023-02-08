/**
  @author: wangyingjie
  @since: 2023/2/8
  @desc: //TODO
**/

package K个一组翻转链表

import (
	SingleList "algorithm/链表/单链表"
	"fmt"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	list := SingleList.NewSingleList()
	list.AddNode(1)
	list.AddNode(2)
	list.AddNode(3)
	list.AddNode(4)
	list.AddNode(5)
	list.AddNode(6)

	fmt.Print("before: ")
	list.Traverse()
	fmt.Print("\n")

	list.Head = reverseKGroup(list.Head, 2)
	list.Traverse()
}
