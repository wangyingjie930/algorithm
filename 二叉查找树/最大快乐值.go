package BST

import (
	"algorithm/二叉查找树/Tree"
	"math"
)

func MaxHappy(root *Tree.Node) int {
	come, nocome := processMaxHappy(root)
	return int(math.Max(float64(come), float64(nocome)))
}

/**
该函数返回的是这个节点来和不来时的快乐值
*/
func processMaxHappy(node *Tree.Node) (int, int) {
	//如果上级来的话, 下级就不来: 父节点快乐值 + 子树(子树根节点不来)的快乐值 + 子树....
	//如果上级不来的话, 下级可来可不来: 0 + MAX(子树根节点不来, 子树根节点来)  + 子树....
	//这里图方便就用二叉树演示了

	if node == nil {
		return 0, 0
	}
	leftCome, leftNoCome := processMaxHappy(node.Left)
	rightCome, rightNoCome := processMaxHappy(node.Right)

	come := node.Value + leftNoCome + rightNoCome
	noCome := 0 + int(math.Max(float64(leftNoCome), float64(leftCome))) +
		int(math.Max(float64(rightNoCome), float64(rightCome)))
	return come, noCome
}
