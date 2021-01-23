package main

import "fmt"

func main() {
	fmt.Println("---------func----------")
	str := longestPalindrome2("ababababad")
	fmt.Println("func:str:", str)
}
// abac
func longestPalindrome2(s string) string {
	if len(s) < 2 { // 肯定是回文，直接返回
		return s
	}

	// 最长回文的首字符索引，和最长回文的长度
	begin, maxLen := 0, 1

	// 在 for 循环中
	// b 代表回文的**首**字符索引号，
	// e 代表回文的**尾**字符索引号，
	// i 代表回文`正中间段`首字符的索引号
	// 在每一次for循环中
	// 先从i开始，利用`正中间段`所有字符相同的特性，让b，e分别指向`正中间段`的首尾
	// 再从`正中间段`向两边扩张，让b，e分别指向此`正中间段`为中心的最长回文的首尾
	for i := 0; i < len(s); { // 以s[i]为`正中间段`首字符开始寻找最长回文。
		if len(s)-i <= maxLen/2 {
			// 因为i是回文`正中间段`首字符的索引号
			// 假设此时能找到的最长回文的长度为l, 则，l <= (len(s)-i)*2 - 1
			// 如果，len(s)-i <= maxLen/2 成立
			// 则，l <= maxLen - 1
			// 则，l < maxLen
			// 所以，无需再找下去了。
			break
		}

		b, e := i, i
		for e < len(s)-1 && s[e+1] == s[e] {
			e++
			// 循环结束后，s[b:e+1]是一串相同的字符串
		}
		// 下一个回文的`正中间段`的首字符只会是s[e+1]
		// 为下一次循环做准备
		fmt.Println("b", b)
		i = e + 1

		// abaabad
		for e < len(s)-1 && b > 0 && s[e+1] == s[b-1] {
			e++
			b--
			// 循环结束后，s[b:e+1]是这次能找到的最长回文。
		}
		// abac
		newLen := e + 1 - b

		// 创新记录的话，就更新记录
		if newLen > maxLen {
			begin = b
			maxLen = newLen
		}
		fmt.Println("len", begin, maxLen)
	}
	return s[begin : begin+maxLen]

}
//1.选中中心点
//2.先处理左右与中心点相同情况
//3.处理左右与中心点不同，但左右相同情况
//4.记住最长值
//5.循环

func longestPalindrome(s string) string {
	maxLen := 0
	maxLenStart := 0
	for i := 0; i < len(s); i++ {
		// 情况1：从单字符向外扩展
		start1, len1 := maxlen(s, i, i)
		fmt.Println(start1, len1)
		// 情况2：从双字符向外扩展
		start2, len2 := maxlen(s, i, i+1)

		// 取较大的
		start, theMax := start1, len1
		if len2 > len1 {
			start, theMax = start2, len2
		}

		// 如果是最大的，保存长度和开始位置
		if theMax > maxLen {
			maxLen = theMax
			maxLenStart = start
		}
	}

	// 截取
	return s[maxLenStart : maxLenStart+maxLen]
}

func maxlen(s string, start, end int) (int, int) {
	// 中心位置扩展，两边不同则停止，停止后两边位置各多扩展了一次
	for start >= 0 && end < len(s) && s[start] == s[end] {
		start--
		end++
	}
	// 开始位置= start+1, 长度=(end-1) - (start+1) + 1
	return start + 1, end - start - 1
}

//链接：https://leetcode-cn.com/problems/longest-palindromic-substring/solution/zhong-xin-kuo-zhan-he-dong-tai-gui-hua-suan-fa-qin/

