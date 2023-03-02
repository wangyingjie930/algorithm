package practice

import (
	"fmt"
	"math/rand"
	"testing"
)

var arr = []int{92, 4, 2, 49, 47, 32, 90, 95, 35, 85, 41, 84, 82, 48, 6, 94, 69, 20, 50, 87, 79, 28,
	67, 15, 36, 55, 75, 93, 91, 56, 8, 12, 42, 16, 66, 59, 57, 5, 33, 1, 54, 97, 80, 77, 65, 52, 29, 81, 14, 13,
	86, 64, 98, 68, 46, 58, 76, 70, 11, 96, 60, 31, 34, 44, 30, 22, 61, 24, 43, 26, 62, 21, 74, 83, 39, 72, 73, 88, 45, 27, 0, 10, 89,
	51, 18, 37, 17, 3, 40, 9, 19, 63, 71, 38, 53, 78, 7, 23, 99, 25}

func TestMajorityElement(t *testing.T) {
	arr := []int{1, 2, 3, 2, 2, 2, 5, 4, 2, 5, 5, 5, 5, 5, 6, 5, 8}
	fmt.Println(MajorityElement(arr))
	//fmt.Print(arr[0:1], arr[1:])
}

func TestFindMedianSortedArrays(t *testing.T) {
	fmt.Println(FindMedianSortedArrays([]int{}, []int{2, 4}))
}

func TestMissingNumber(t *testing.T) {
	/*h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}*/
	fmt.Println(MissingNumber([]int{3, 0, 1}))
}

func TestMinWindow(t *testing.T) {
	fmt.Println(MinWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(MinWindow("aa", "aa"))
	fmt.Println(MinWindow("a", "b"))
}

func TestPrintArr(t *testing.T) {
	fmt.Println(SolveNQueens(9))
}

func TestLongestPalindrome(t *testing.T) {
	fmt.Println(MaxLcpsLength1("babad"))
	fmt.Println(LongestPalindrome("babad"))
	fmt.Println(1&1, 0&1, 2&1)
}

func TestNumIslands(t *testing.T) {
	fmt.Println(NumIslands([][]byte{
		{1, 0, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}))
	/*
		[["1","1","1"],["0","1","0"],["1","1","1"]]	*/
	//[["1","0","1","1","1"],["1","0","1","0","1"],["1","1","1","0","1"]]
	fmt.Println(NumIslands([][]byte{
		{1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1},
		{1, 1, 1, 0, 1},
	}))
}

func TestTwoMi(t *testing.T) {
	fmt.Println(TwoMi(1024))
	fmt.Println(IsPowerOfFour(16), 0x55)
}

func TestMinSubArrayLen(t *testing.T) {
	fmt.Println(arr, MinSubArrayLen(110, arr))
}

func TestMinSubArrayLenDp(t *testing.T) {
	fmt.Println(MinSubArrayLenDp(63542, rand.Perm(10000)))
}

func TestMinSubArrayLenDynamic(t *testing.T) {
	fmt.Println(MinSubArrayLenDynamic(63542, rand.Perm(10000)))
}

func TestSmallestK(t *testing.T) {
	fmt.Println(SmallestK([]int{-1, 3, 4, 0, -1, 2, 5, -6, 3, -2, -2, 0, -1, -1, 2}, 4))
}

func TestLongestValidParentheses(t *testing.T) {
	fmt.Println(LongestValidParentheses(")()())()"))
}

func TestNumDecodings(t *testing.T) {
	fmt.Println(NumDecodings("100"))
	fmt.Println(NumDecodings1("111111111111"))
}

func TestFindMinMoves(t *testing.T) {
	fmt.Println(FindMinMoves([]int{9, 1, 8, 8, 9}))
	fmt.Println(FindMinMoves([]int{0, 3, 0}))
}

func TestBeer(t *testing.T) {
	fmt.Println(Beer1(10))
}
