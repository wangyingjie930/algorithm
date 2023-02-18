package BST

import (
	BST2 "algorithm/二叉查找树"
	"fmt"
	"strings"
	"testing"
)

var tree = newBST()

func newBST() *BST {
	tree := NewBST()
	arr := []int{385, 823, 441, 805, 258, 765, 868, 83, 850, 379, 812, 113, 571, 422, 123, 189, 364, 587, 637,
		399, 629, 244, 752, 529, 822, 599, 788, 514, 706, 796, 611, 943, 511, 788, 529, 900, 438, 808, 559, 335, 158,
		902, 39, 147, 397, 529, 176, 338, 530, 75, 741, 449, 951, 304, 221, 574, 439, 500, 412, 989, 666, 436, 867, 701,
		793, 884, 759, 679, 736, 375, 687, 497, 44, 365, 915, 410, 244, 247, 840, 428, 786, 236, 807, 989, 700, 218, 758, 912, 897, 749, 656,
		650, 140, 150, 182, 80, 213, 450, 56, 240}
	arr1 := []int{92, 4, 2, 49, 47, 32, 90, 95, 35, 85, 41, 84, 82, 48, 6, 94, 69, 20, 50, 87, 79, 28,
		67, 15, 36, 55, 75, 93, 91, 56, 8, 12, 42, 16, 66, 59, 57, 5, 33, 1, 54, 97, 80, 77, 65, 52, 29, 81, 14, 13,
		86, 64, 98, 68, 46, 58, 76, 70, 11, 96, 60, 31, 34, 44, 30, 22, 61, 24, 43, 26, 62, 21, 74, 83, 39, 72, 73, 88, 45, 27, 0, 10, 89,
		51, 18, 37, 17, 3, 40, 9, 19, 63, 71, 38, 53, 78, 7, 23, 99, 25}

	for k, v := range arr {
		tree.Insert(arr1[k], v)
	}

	return &tree
}

func TestBST_PreOrder(t *testing.T) {
	tree.PreOrder()
}

func TestBST_InOrder(t *testing.T) {
	tree.InOrder()
}

func TestBST_PostOrder(t *testing.T) {
	tree.PostOrder()
}

func TestBST_Maximum(t *testing.T) {
	fmt.Println(tree.Maximum())
}

func TestBST_Minimum(t *testing.T) {
	fmt.Println(tree.Minimum())
}

func TestBST_RemoveMax(t *testing.T) {
	tree.RemoveMax()
	tree.InOrder()
}

func TestBST_RemoveMin(t *testing.T) {
	tree.RemoveMin()
	tree.InOrder()
}

func TestBST_Remove(t *testing.T) {
	tree.Remove(75)
	tree.InOrder()
}

func TestLowestAncestor(t *testing.T) {
	node := BST2.LowestAncestor(tree.Root, tree.Root.Left, tree.Root.Right.Right.Left)
	fmt.Printf("LCA: %+v\nroot: %+v\n", node, tree.Root)

	node = BST2.LowestAncestor(tree.Root, tree.Root.Left, tree.Root.Right.Right.Left)
	fmt.Printf("LCA: %+v\nroot: %+v\n", node, tree.Root)

	node = BST2.LowestAncestor(tree.Root, tree.Root.Left, tree.Root.Left.Right)
	fmt.Printf("LCA: %+v\nroot: %+v\n", node, tree.Root.Left)

	node = BST2.LowestAncestor(tree.Root, tree.Root.Left.Left, tree.Root.Left.Right.Right)
	fmt.Printf("LCA: %+v\nroot: %+v\n", node, tree.Root.Left)
}

func TestIsBalance(t *testing.T) {
	tree := NewBST()
	tree.Insert(3, 3)
	tree.Insert(1, 1)
	tree.Insert(2, 2)
	tree.Insert(4, 4)
	fmt.Print(BST2.IsBalance(tree.Root))
}

func TestIsBCT(t *testing.T) {
	tree := NewBST()
	tree.Insert(2, 2)
	tree.Insert(3, 3)
	tree.Insert(1, 1)
	fmt.Print(BST2.IsBCT(tree.Root))
}

func TestLevelOrder(t *testing.T) {
	fmt.Println(BST2.LevelOrder(tree.Root))
}

func TestSerialize(t *testing.T) {
	serial := BST2.Serialize(tree.Root)
	fmt.Println(serial)
	tree.Root = BST2.UnSerialize(serial)
	tree.PreOrder()
}

func TestString(t *testing.T) {
	fmt.Println(strings.Index("92,4,2,1,0", ","))
}

func TestMaxHappy(t *testing.T) {
	fmt.Println(BST2.MaxHappy(tree.Root))
}
