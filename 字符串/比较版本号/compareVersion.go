/**
  @author: wangyingjie
  @since: 2023/4/3
  @desc: https://leetcode.cn/problems/compare-version-numbers/description/
**/

package 比较版本号

import "strconv"

/*
使用的是字符串切割
func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	for i1, i2 := 0, 0; i1 < len(v1) || i2 < len(v2); {
		t1 := 0
		if i1 < len(v1) {
			t1, _ = strconv.Atoi(v1[i1])
		}
		t2 := 0
		if i2 < len(v2) {
			t2, _ = strconv.Atoi(v2[i2])
		}

		if t1 > t2 {
			return 1
		} else if t1 < t2 {
			return -1
		}

		i1++
		i2++
	}
	return 0
}*/

// compareVersion
//  @Description: 使用双指针
//  @param version1
//  @param version2
//  @return int
func compareVersion(version1 string, version2 string) int {
	s1, s2 := 0, 0
	f1, f2 := s1, s2
	for f1 < len(version1) || f2 < len(version2) {
		for f1 < len(version1) && version1[f1] != '.' {
			f1++
		}
		t1 := getInt(version1, s1, f1)
		s1 = f1 + 1
		f1 = s1

		for f2 < len(version2) && version2[f2] != '.' {
			f2++
		}
		t2 := getInt(version2, s2, f2)
		s2 = f2 + 1
		f2 = s2

		if t1 > t2 {
			return 1
		} else if t1 < t2 {
			return -1
		}
	}
	return 0
}

func getInt(s string, start, end int) int {
	if start >= len(s) {
		return 0
	}
	ret, _ := strconv.Atoi(s[start:end])
	return ret
}
