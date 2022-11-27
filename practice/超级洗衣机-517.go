package practice

import "math"

func FindMinMoves(machines []int) int {
	//1. 左: 多出x件衣服 右多y件衣服, 中间可以同时接收衣服, 则 中 = MAX(x, y),
	//2. 左: 少x件衣服 少y件衣服, 中间不能同时往左右两边分衣服, 则 中 = x + y,
	//3. 左: 少x件, 右: 多y件[次数为Max(x, y), 若y>x, 则右→中→左x次后, 右→中 y-x次, 若....]
	//4. 左: 多x件, 右: 少y件, 同理 Max(x, y)
	sum := 0
	for _, v := range machines {
		sum += v
	}
	if sum % len(machines) != 0 {
		return -1
	}
	avg := sum / len(machines)

	tmp := 0
	maxNum := 0
	for k, v := range machines {
		left := tmp - avg * k
		right := (sum - v - tmp) - avg * (len(machines) - k - 1)

		if left < 0 && right < 0 {
			if maxNum < int(math.Abs(float64(left)) + math.Abs(float64(right))) {
				maxNum = int(math.Abs(float64(left)) + math.Abs(float64(right)))
			}
		}else {
			bigger := math.Abs(float64(left))
			if bigger < math.Abs(float64(right)) {
				bigger = math.Abs(float64(right))
			}
			if int(bigger) > maxNum {
				maxNum = int(bigger)
			}
		}
		tmp += v
	}
	return maxNum
}
