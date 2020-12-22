package main

import "fmt"

//请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。
//给定一个string，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。
//解题思路
//翻转字符串其实是将一个字符串以中间字符为轴，前后翻转，即将str[len]赋值给str[0],将str[0] 赋值 str[len]。
func main() {
	var str = "大溪谷abcdefgh"
	fmt.Println(reverStr(str))
}
func reverStr(s string) (string, bool) {
	str := []rune(s)
	l := len(str)
	if l > 5000 {
		return s, false
	}
	for i := 0; i < l/2; i++ {
		str[i], str[l-i-1] = str[l-i-1], str[i]
	}
	return string(str), true
}
