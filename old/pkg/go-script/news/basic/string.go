/*
@Time : 2019/5/3 5:49 PM 
@Author : Tenlu
@File : string
@Software: GoLand
*/
package main

import "fmt"

func main() {
	// 默认默认UTF-8编码,中文占3个字节,unicode占2个字节
	
	//// Go语言的字符串的字节使用UTF-8编码标识Unicode文本。
	////Go语言的字符串的字节使用UTF-8编码标识Unicode文本,这样Golang统一使用UTF-8编码,中文乱码问题不会再困扰程序员。
	////字符串一旦赋值了,字符串就不能修改了:在Go中字符串是不可变的。
	// golang 遍历中文字符串,	遍历是按照字节遍历，因此如果有中文等非英文字符，就会出现乱码
	str := "hi,北京"
	fmt.Printf("len:%d\n", len(str)) // 9
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}

	fmt.Println("---按照utf8编码--")
	for i, ch := range str {
		fmt.Printf("str[%d]=%c\n", i, ch)
	}
	// 结果如下:str[3]~str[5]三个字符组合了一个京字,显示下标直接跳过了4,5
	//str[0]=h
	//str[1]=i
	//str[2]=,
	//str[3]=北
	//str[6]=京
	fmt.Println("----切片方式遍历中文---")
	strArr := []rune(str)
	for i, c := range strArr {
		fmt.Printf("str[%d]=%c\n", i, c)
	}
	//len:9
	//str[0]=h
	//str[1]=i
	//str[2]=,
	//str[3]=å
	//str[4]=å*
	//str[5]=å*
	//str[6]=ä
	//str[7]=º
	//str[8]=¬
	//---按照utf8编码--
	//str[0]=h
	//str[1]=i
	//str[2]=,
	//str[3]=北
	//str[6]=京
	//----切片方式遍历中文---
	//	str[0]=h
	//str[1]=i
	//str[2]=,
	//str[3]=北
	//str[4]=京
	strstr()
	fmt.Printf("%s\n",reverse("hello world"))

}

func strstr()  {
	println("-----字符串截取,拼接-------")
	str:="hello world"
	hello:=str[:5]
	world:=str[5:]
	fmt.Printf("%s\n",hello+world)
	//rune用于表示每个Unicode码点，目前只使用了21个bit位
	for _, c :=range []rune(str) {
		fmt.Printf("%c",c)
	}
	// 字符串 "hello,world"和下面的字节数组底层一致
	var data = [...]byte{'h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd'}
	for _,d:=range data {
		fmt.Printf("%s",string(d))
	}


	fmt.Println()
}

// 字符串翻转
func reverse(str string) string  {
	fmt.Println("----字符串翻转--------")
	var result string
	strlen:=len(str)
	for i:=0;i<strlen;i++{
		 // %c 输出单个字符
		 result = result+fmt.Sprintf("%c",str[strlen-1-i])
	}
	return result
}






