package main

import "fmt"

// https://leetcode-cn.com/problems/implement-strstr/
func strStr(haystack string, needle string) int {
	needleLength := len(needle) // 提取，循环中不计算长度
	if needleLength == 0 { // 加速，可以去掉
		return 0
	}
	haystackLength := len(haystack) // 提取，循环中不计算长度
	for k := range haystack {
		if k+needleLength <= haystackLength {
			if haystack[k:k+needleLength] == needle {
				return k
			}
		} else { // 加速，可以去掉
			return -1
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("a", "a"))
}
