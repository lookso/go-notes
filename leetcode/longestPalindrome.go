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
	str := longestPalindrome("bbbabcba")
	fmt.Println("func1:str:",str)
	fmt.Println("---------func2----------")
	str2 := longestPalindrome2("bbbabcba")
	fmt.Println("func2:str2:",str2)

	fmt.Println("---------func3----------")
	str3 := longestPalindrome3("bbbabcba")
	fmt.Println("func3:str3:",str3)
}

// check whether a string is palindromic
func isPalindromic(s string) bool {
	mid := len(s) / 2
	last := len(s) - 1
	for i := 0; i < mid; i++ {
		if s[i] != s[last-i] {
			return false
		}
	}
	return true
}
// 穷举法 时间复杂度On的三次方
// longest palindromic substring
func longestPalindrome(s string) string {
	last := len(s) - 1
	longest := string(s[0])  //a
	fmt.Println(longest)
	for i := 0; i < last; i++ {
		for j := i + 1; j <= last; j++ {
			if isPalindromic(s[i:j+1]) && j+1-i > len(longest) {
				longest = s[i : j+1]
			}
		}
	}
	return longest
}

// 对二维状态数组𝑐采取“之”字形赋值，增加gap来扩大子串长度；有两层for循环，故复杂度为𝑂(𝑛2次方).
func longestPalindrome2(s string) string {
	length := len(s)
	// longest palindromic substring
	longest := string(s[0])
	// 2-D array initialization
	c := make([][]bool, length)
	for i := range c {
		c[i] = make([]bool, length)
	}
	for gap := 0; gap < length; gap++ {
		for i := 0; i < length - gap; i++ {
			j := i + gap
			if s[i] == s[j] && (j - i <= 2 || c[i + 1][j - 1]) {
				c[i][j] = true
				// find the longest
				if j + 1 - i > len(longest) {
					longest = s[i:j + 1]
				}
			}
		}
	}
	return longest
}


// given a center, find the longest palindromic substring
func l2rHelper(s string, mid int) string {
	l := mid - 1; r := mid + 1
	length := len(s)
	for r < length && s[r] == s[mid] {
		r++
	}
	for l >= 0 && r < length && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1:r]
}
// 对于给定中心，找出最长回文子串的复杂度为𝑂(𝑛)；总的复杂度为𝑂(𝑛2次方).
func longestPalindrome3(s string) string {
	length := len(s)
	longest := string(s[0])
	for i := 0; i < length -1; i++  {
		if len(l2rHelper(s, i)) > len(longest) {
			longest = l2rHelper(s, i)
		}
	}
	return longest
}