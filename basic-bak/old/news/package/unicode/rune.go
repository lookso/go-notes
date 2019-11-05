/*
@Time : 2019-05-21 22:21 
@Author : Tenlu
@File : rune
@Software: GoLand
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main()  {
	runeFunc()
}
func runeFunc()  {
	str := "hello 世界"
	fmt.Println(len([]rune(str)))       // 8
	fmt.Println(utf8.RuneCountInString(str)) //8
	fmt.Println(len([]byte(str)))  // 12
	fmt.Println(len(str))       // 12    由此可见string的底层实现就是byte

	fmt.Println(utf8.RuneLen('a')) // 1
	fmt.Println(utf8.RuneLen('界')) // 3
	// 返回p中的utf-8编码的码值的个数。错误或者不完整的编码会被视为宽度1字节的单个码值。
	buf := []byte("Hello, 世界")
	fmt.Println("bytes =", len(buf))
	fmt.Println("runes =", utf8.RuneCount(buf))
	fmt.Println("runes =",len([]rune("Hello, 世界")))

}
