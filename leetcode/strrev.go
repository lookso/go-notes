/*
@Time : 2019-10-28 22:25
@Author : Tenlu
@File : strrev
@Software: GoLand
*/
package main

import "fmt"

// 字符串反转
func main() {
	fmt.Println(strrev("nba"))
}
func strrev(str string) string {

	strRune := []rune(str)
	strLen := len(strRune)
	newStrRune := make([]rune, 0)
	for i := strLen - 1; i >= 0; i-- {
		newStrRune = append(newStrRune, strRune[i])
	}
	return string(newStrRune)

	// 或者第二种
	for i := 0; i < strLen/2; i++ {
		tmp := strRune[i]
		strRune[i] = strRune[strLen-1-i]
		strRune[strLen-1-i] = tmp
	}
	return string(strRune)

}
