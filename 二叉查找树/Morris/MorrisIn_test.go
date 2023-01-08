/**
  @author: wangyingjie
  @since: 2022/11/27
  @desc: //TODO
**/

package Morris

import (
	BST "algorithm/二叉查找树"
	"fmt"
	"testing"
)

var tree = newBST()

func newBST() BST.Node {
	node4:=BST.Node{Key: 4, Value: 4}
	node5:=BST.Node{Key: 5, Value: 5}
	node6:=BST.Node{Key: 6, Value: 6}
	node7:=BST.Node{Key: 7, Value: 7}

	node2 := BST.Node{Key: 2, Value: 2, Left: &node4, Right: &node5}
	node3 := BST.Node{Key: 3, Value: 3, Left: &node6, Right: &node7}

	return BST.Node{Key: 1, Value: 1, Left: &node2, Right: &node3}
}

func Test_morrisIn(t *testing.T) {
	fmt.Println()
	morrisIn(&tree)
}
