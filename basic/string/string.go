/*
@Time : 2019-10-30 21:53 
@Author : Tenlu
@File : string
@Software: GoLand
*/
package main

import "fmt"

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
	for kk,vv:=range strRune{
		fmt.Println(kk,vv)
	}
	fmt.Println("-----byte----")
	for kkk,vvv:=range strByte{
		fmt.Println(kkk,vvv)
	}
}
