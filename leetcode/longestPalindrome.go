/*
@Time : 2019-10-28 18:36 
@Author : Tenlu
@File : longestPalindrome
@Software: GoLand
*/
package main

import "fmt"

// 求出最长回文子串

func main() {
	fmt.Println("---------func----------")
	str := longestPalindrome("abat")
	fmt.Println("func:str:", str)

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
		fmt.Println(start1,len1)
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