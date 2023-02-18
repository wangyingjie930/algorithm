/**
  @author: wangyingjie
  @since: 2023/2/18
  @desc: //TODO
**/

package Tree

type Node struct {
	Key   int
	Value int
	Left  *Node
	Right *Node
}

func (node *Node) GetLeft() (TreeNode, bool) {
	return node.Left, node.Left == nil
}

func (node *Node) GetRight() (TreeNode, bool) {
	return node.Right, node.Right == nil
}

func (node *Node) GetKey() int {
	return node.Key
}

func (node *Node) GetValue() int {
	return node.Value
}

// NewTreeByNums
//  @Description: 通过数组生成普通的二叉树
//  @param nums
//  @param i
//  @return *Node
func NewTreeByNums(nums []int, i int) *Node {
	if i < 0 || len(nums) <= i || nums[i] == -1 {
		return nil
	}
	node := &Node{Value: nums[i], Key: nums[i]}
	node.Left = NewTreeByNums(nums, 2*i+1)
	node.Right = NewTreeByNums(nums, 2*i+2)
	return node
}
