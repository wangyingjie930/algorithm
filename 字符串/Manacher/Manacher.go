package Manacher

import "math"

// 对长度为偶数的字符串进行#的扩充, 使它成为奇数的长度
func manacherString(subject string) string {
	ms := ""
	index := 0
	for i := 0; i < len(subject)*2+1; i++ {
		if i%2 == 0 {
			ms = ms + "#"
		}else {
			ms = ms + string(subject[index])
			index ++
		}
	}
	return ms
}

func maxLcpsLength(subject string) int  {
	ms := manacherString(subject)

	//只要有字符的半径比当前的半径还要往右, 那么就进行记录
	R := -1
	//R更新时C一定更新, 代表的是R更新时候的该字符的索引
	C := -1

	//每个字符对应最大的半径是多少
	radiusArr := make([]int, len(ms))
	maxRadius := 0
	for i := range ms {
		if R < i {
			radiusArr[i] = 0
		}else {
			//这里就是包含两种情况, radiusArr[i]为至少不用验证的区域
			radiusArr[i] = int(math.Min(float64(R -i), float64(radiusArr[2 * C - i])))
		}
		//进行暴力扩
		for j := radiusArr[i] + 1; i - j >= 0 && i + j < len(ms) && ms[i + j] == ms[i - j]; j++ {
			radiusArr[i]++
		}
		//获得最右的R和C
		if radiusArr[i] + i > R {
			R = radiusArr[i] + i
			C = i
		}
		//得到最大的半径
		if radiusArr[i] > maxRadius {
			maxRadius = radiusArr[i]
		}
	}
	return maxRadius / 2
}


