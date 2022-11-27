package 暴力破解

import (
	"github.com/syyongx/php2go"
)

func BF(needle string, search string) int {
	for i := 0; i < len(needle); i ++ {
		k := i
		j := 0
		for ; j < len(search); {
			if needle[k] == search[j] {
				j ++
				k ++
			}else {
				break
			}
		}
		if j == len(search) {
			return i
		}
	}
	return -1
}

func RK(needle string, search string) int {
	l := len(search)
	for k :=0; k + l <= len(needle); k ++{
		if php2go.Sha1(needle[k:l+k]) == php2go.Sha1(search) {
			return k
		}
	}
	return -1
}
