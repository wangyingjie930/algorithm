package main

import (
	"fmt"
	"math/rand"
	"time"
)

func reservoirSampling(stream []int, m int) []int {
	rand.Seed(time.Now().UnixNano())

	reservoir := make([]int, m)
	for i := 0; i < m; i++ {
		reservoir[i] = stream[i]
	}
	for i := m; i < len(stream); i++ {
		j := rand.Intn(i + 1)
		if j < m {
			reservoir[j] = stream[i]
		}
	}
	return reservoir
}

func main() {
	stream := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	m := 3
	reservoir := reservoirSampling(stream, m)
	fmt.Println("Reservoir Sampling Results:")
	fmt.Println(reservoir)
}
