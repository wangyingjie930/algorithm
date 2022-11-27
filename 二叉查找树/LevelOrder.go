package BST

/**
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

    3
   / \
  9  20
    /  \
   15   7


[
  [3],
  [9,20],
  [15,7]
]
 */
func LevelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*Node{root}
	var ret [][]int
	var tmp []int
	for len(queue) > 0 {
		q := queue
		queue = nil
		tmp = []int{}
		for _, node := range q {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			tmp = append(tmp, node.Key)
		}
		ret = append(ret, tmp)
	}
	return ret
}

/**
给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

[
  [3],
  [20,9],
  [15,7]
]
 */
func ZigzagLevelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*Node{root}
	var ret [][]int
	level := 0
	var tmp []int

	for len(queue) > 0 {
		q := queue
		queue = nil
		//目的是为了将下一层的所有节点直接放到队列里面来
		for _, node := range q {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		tmp = []int{}
		if level % 2 == 0 {
			//当前层数为偶数, 从左向右打印
			for i := 0; i < len(q); i++ {
				tmp = append(tmp, q[i].Key)
			}
		}else {
			//当前层数为奇数, 从右向左打印
			for i := len(q) - 1; i >= 0; i-- {
				tmp = append(tmp, q[i].Key)
			}
		}
		ret = append(ret, tmp)
		level ++
	}
	return ret
}