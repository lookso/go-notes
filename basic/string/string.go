/*
@Time : 2019-10-30 21:53 
@Author : Tenlu
@File : string
@Software: GoLand
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "nb中国"
	strRune := []rune(str)
	strByte := []byte(str)
	fmt.Println("len(str)", len(str))         // 8 // 默认是按utf-8编码,一个utf-8 占3个字节
	fmt.Println("len(strRune)", len(strRune)) // 4
	fmt.Println("len(strByte)", len(strByte)) // 8
	for k, ch := range str {
		fmt.Printf("str[%d]=%c\n", k, ch)
		fmt.Println(k, ch)
	}
	fmt.Println("-----rune----")
	for kk, vv := range strRune {
		fmt.Println(kk, vv)
	}
	fmt.Println("-----byte----")
	for kkk, vvv := range strByte {
		fmt.Println(kkk, vvv)
	}
	// 双引号内部变量会解析,反引号则不解析,注意不是单引号
	// 反引号串中的内容总被认为是普通字符，因此反引号中的内容不会被转义效率更高
	str1 := "nba\n"        // 会被换行
	str2 := `nba\n`        // 上面原生字符串中的 \n 会被原样输出。
	fmt.Println(len(str1)) // 4
	fmt.Println(len(str2)) // 5

	var hello = "helloworld"
	newString := strings.Replace(hello, hello[strings.Index(hello, "w"):], "python", 1)
	fmt.Println(newString) // hellopython

	helloByte := []byte(hello) //len = 10
	helloByte[0] = 'w'         // 只能有一个字节元素,切记
	helloByte[1] = 'e'
	helloByte[2] = 'l'
	helloByte[3] = 'c'
	helloByte[4] = 'o'
	helloByte[5] = 'm'
	helloByte[6] = 'e'

	fmt.Println(helloByte)         // [119 101 108 99 111 109 101 114 108 100]
	fmt.Println(string(helloByte)) // welcomerld
}


