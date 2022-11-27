package practice

func TwoMi(a int) bool {
	b := a - 1
	if a & b == 0 {
		return true
	}
	return false
}

func IsPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	if n&(n-1) == 0 && n&0x55555555 != 0 {
		return true
	}
	return false
}
