/*
@Time : 2019-05-21 21:58 
@Author : Tenlu
@File : unicode
@Software: GoLand
*/
package main

import (
	"fmt"
	"unicode/utf8"
	"unicode"
)

func main()  {
	valid := 'a'
	// 将rune理解为 一个 可以表示unicode 编码的值int 的值，称为码点（code point)
	// byte()和rune()一个是uint8、一个是uint32。就这么简单~。
	invalid := rune(0xfffffaa)
	fmt.Println(utf8.ValidRune(valid))
	fmt.Println(utf8.ValidRune(invalid))
	fmt.Println(unicode.ToLower(rune(invalid)))

	
}
