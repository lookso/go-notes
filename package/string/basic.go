/*
@Time : 2019/2/28 10:55 AM 
@Author : Tenlu
@File : basic
@Software: GoLand
*/
package main

import (
	"fmt"
	"strings"
)

//strings包实现了用于操作字符的简单函数。
func main() {
	// 判断两个utf-8编码字符串（将unicode大写、小写、标题三种格式字符视为相同）是否相同。不区分大小写
	fmt.Println(strings.EqualFold("Go", "go"))
	// == 是区分字符串大小写的
	if "go" == "Go" {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	fmt.Println(strings.Compare("o", "Go"))
	// 判断s是否有前缀字符串prefix。
	fmt.Println(strings.HasPrefix("helloworld", "hello"))
	// 判断s是否有后缀字符串suffix。
	fmt.Println(strings.HasSuffix("helloworld", "world"))
	// 判断字符串s是否包含子串substr
	fmt.Println(strings.Contains("seafood", "foo"))
	// 判断字符串s是否包含字符串chars中的任一字符
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	// 子串sep在字符串s中第一次出现的位置，不存在则返回-1。
	fmt.Println(strings.Index("failure", "re"))

	//字符c在s中第一次出现的位置，不存在则返回-1。
	fmt.Println(strings.IndexByte("failure", 'u')) // 此处单引号作为byte类型的值,双引号就是字符串 // 4
	// 将字符串全部转换为小写
	fmt.Println(strings.ToLower("HeLLo")) //hello
	// 将字符串全部转换为小写
	fmt.Println(strings.ToUpper("heLLo")) //HELLO
	// 返回字符串s中有几个不重复的sep子串。如果子字符串是空，就返回字符数+1
	fmt.Println(strings.Count("hello中国", "") - 1)
	// 在 Golang中,如果字符串中出现中文字符不能直接调用 len 函数来统计字符串字符长度，
	// 这是因为在 Go 中,字符串是以 UTF-8 为格式进行存储的,在字符串上调用 len 函数,取得的是字符串包含的 byte 的个数。
	fmt.Println(len("hello中国"))                  // 一个utf8中文3个字节
	fmt.Println(strings.Count("cheesela", "e"))  // 3
	fmt.Println(strings.Count("cheesela", "ee")) // 1
	// 将一个字符串数组连接为一个字符串,之间用sep来分隔。类似php的implode()
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ",")) // foo,bar,baz

	fmt.Printf("%q\n", strings.Split("a,b,c", ",")) //["a" "b" "c"] 类似php的explode()函数

	// 返回将s前后端所有cutset包含的utf-8码值都去掉的字符串。
	fmt.Println(strings.Count(strings.Trim(" hello jack,your are god ", " "), "") - 1)
	fmt.Println(strings.TrimLeft(" hello", " "))
	// NewReader创建一个从s读取数据的Reader。本函数类似bytes.NewBufferString，但是更有效率，且为只读的。
	fmt.Println("-------------------")
	reader := strings.NewReader("hello world")
	b := make([]byte, 8)
	n, _ := reader.Read(b)     // 8
	fmt.Println(string(b[:n])) // hello world 这个时候没有读到要读d
	reader.Seek(2, 1)          //这个时候我们在读取位置 偏移2
	reader.UnreadByte()        //然后向前偏移1字节
	n1, _ := reader.Read(b)
	fmt.Println(string(b[:n1])) //ld

}
