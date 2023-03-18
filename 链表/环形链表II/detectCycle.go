/**
  @author: wangyingjie
  @since: 2023/3/18
  @desc: https://leetcode.cn/problems/linked-list-cycle-ii/description/
**/

package 环形链表II

import List "algorithm/链表/单链表"

func detectCycle(head *List.Node) *List.Node {
	slow, quick := head, head

	for slow != nil {
		slow = slow.Next
		for i := 0; i < 2; i++ {
			//快指针这块还是搞个循环比较方便, 不要搞quick.Next.Next, 这样判断起来很麻烦
			if quick != nil {
				quick = quick.Next
			} else {
				return nil
			}
		}

		if slow == quick {
			break
		}
	}

	slow = head
	for slow != quick {
		//slow回到开头, 大家都走一步
		slow = slow.Next
		quick = quick.Next
	}

	return slow
}
