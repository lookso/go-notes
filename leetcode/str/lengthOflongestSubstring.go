package main

import "fmt"

// 最长不含重复字符的子字符串
// 请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
func main() {
	fmt.Println(lengthOfLongestSubstring("fabca"))
}

//哈希检测到是否有相同值
//无重复值，当前字符收录入哈希，右指针扩张
//有则，右指针等待扩大边界，左指增长，同时哈希收缩
//重复直到同值左侧耗尽，一轮循环时记录最大长度
// bafrbc
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	left, right := 0, 0
	length := 0
	for ; right < len(s); {
		if _, ok := window[s[right]]; ok {
			//发现了重复
			delete(window, s[left])
			left++
		} else {
			if right-left >= length {
				length = right - left + 1
			}
			fmt.Println(string(s[right]))
			//不存在重复
			window[s[right]] = 1
			right++
		}
	}
	return length
}

//作者：yanbinliu
//链接：https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/solution/jian-zhi-offer-48-zui-chang-bu-han-zhong-fu-zi-13/

//func lengthOfLongestSubstring(s string) int {
//	lastOccured := make(map[byte]int)
//	start := 0
//	maxLength := 0
//	for i, ch := range []byte(s) {
//		lastIndex, ok := lastOccured[ch]
//		if ok && lastIndex >= start {
//			start = lastOccured[ch] + 1
//		}
//		fmt.Println(start)
//		if i-start+1 > maxLength {
//			maxLength = i - start + 1
//		}
//		fmt.Println("maxLength",maxLength)
//		lastOccured[ch] = i
//	}
//	return maxLength
//	//链接：https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/solution/zui-hao-li-jie-de-fang-fa-by-ni-yu-shi-guang-jie-q/
//}

