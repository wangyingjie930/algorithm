package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {
	data := []int{10, 11, 17, 22, 22, 26}
	index, err := binarySearch(data, 22)
	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.Unwrap(err))
		os.Exit(1)
	}
	fmt.Print(index)
}

/** 二分查找 */
func binarySearch(data []int, search int) (int, error) {
	i := 0
	j := len(data)
	mid := int(math.Ceil(float64((i + j) / 2)))
	if j == 1 && data[0] != search {
		return -1, errors.New("找不到")
	}
	if data[mid] == search {
		return mid, nil
	}else if data[mid] > search {
		j = mid
	}else {
		i = mid + 1
		if i >= j {
			i = j - 1
		}
	}
	index, err := binarySearch(data[i:j], search)
	if err != nil {
		//return -1, err
		return -1, fmt.Errorf("w : %w", err)
	}
	return i + index, nil
}
