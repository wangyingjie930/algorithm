package practice

/**
使用堆排序
 */
func FindKthLargest(nums []int, k int) int {
	heap := newHeap(nums)
	for k > 1 {
		heap[1], heap[len(heap) - 1] = heap[len(heap) - 1], heap[1]
		heap = heap[:len(heap) - 1]
		shiftDownMax(heap, 1)
		k --
	}
	return heap[1]
}

func newHeap(nums []int) []int {
	heap := append([]int{0}, nums...)
	for i := (len(heap) - 1) / 2; i > 0; i -- {
		heap = shiftDownMax(heap, i)
	}
	return heap
}

func shiftUpMax(nums []int, i int) []int {
	for j := i; j / 2 > 0 && nums[j / 2] < nums[j]; j = j / 2 {
		nums[j], nums[j / 2] = nums[j / 2], nums[j]
	}
	return nums
}

func shiftDownMax(nums []int, i int) []int {
	for j := i; j < len(nums); {
		max := nums[j]
		maxIndex := j
		if 2 * j < len(nums) && max < nums[2 * j] {
			maxIndex = 2 * j
			max = nums[maxIndex]
		}
		if 2 * j + 1 < len(nums) && max < nums[2 * j + 1] {
			maxIndex = 2 * j + 1
			max = nums[maxIndex]
		}
		if maxIndex == j {
			break
		}
		nums[j], nums[maxIndex] = nums[maxIndex], nums[j]
		j = maxIndex
	}
	return nums
}
