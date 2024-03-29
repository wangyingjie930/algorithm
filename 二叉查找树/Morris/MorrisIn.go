/**
  @author: wangyingjie
  @since: 2022/11/27
  @desc: morris遍历
**/

package Morris

import (
	"algorithm/二叉查找树/Tree"
	"fmt"
)

func morrisIn(root *Tree.Node) {
	head := root
	for head != nil {
		if head.Left != nil {
			//从左树中获得最右的节点
			right := head.Left
			//添加right.Right != head是因为怕因为线索形成无限循环
			for right.Right != nil && right.Right != head {
				right = right.Right
			}
			if right.Right == nil {
				//fmt.Println(head.Value) //前序
				//第一次到最右的节点, 进行线索的创建
				right.Right = head
				//指针指向下一个节点
				head = head.Left
			} else {
				//通过线索创建过, 表示已经通过线索回来过, 不再需要线索, 需要还原回去
				right.Right = nil
				fmt.Println(head.Value) //中序
				//开始向右子树遍历
				head = head.Right
			}
		} else {
			//fmt.Println(head.Value) //前序
			fmt.Println(head.Value) //中序
			//通过线索返回
			head = head.Right
		}
	}
}
