package main

import "fmt"

// 最长不含重复字符的子字符串
// 请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
func main() {
	fmt.Println(lengthOfLongestSubstring("bafrbcd"))
}
func lengthOfLongestSubstring(s string) int {
	lastOccured := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		lastIndex, ok := lastOccured[ch]
		if ok && lastIndex >= start {
			start = lastOccured[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxLength
	//链接：https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/solution/zui-hao-li-jie-de-fang-fa-by-ni-yu-shi-guang-jie-q/
}
