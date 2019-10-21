/*
@Time : 2019-10-20 10:42 
@Author : Tenlu
@File : strrev
@Software: GoLand
*/
package main

import "fmt"

func main() {
	var str = "nba中国"
	fmt.Println(strrev(str))
}

func strrev(str string) string {
	strRune := []rune(str)
	len := len(strRune)
	if len > 5000 {
		return string(str)
	}
	newStr := make([]rune, 10)
	for i := range strRune {
		newStr = append(newStr, strRune[len-1-i])
	}
	return string(newStr)
	// 或者第二种
	for i := 0; i < len/2; i++ {
		tmp := strRune[i]
		strRune[i] = strRune[len-1-i]
		strRune[len-1-i] = tmp
	}
	return string(strRune)
}
