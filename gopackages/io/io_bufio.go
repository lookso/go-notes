/*
@Time : 2019/2/28 10:48 AM 
@Author : Tenlu
@File : basic
@Software: GoLand
*/
package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

//bufio包实现了有缓冲的I/O。
//它包装一个io.Reader或io.Writer接口对象，创建另一个也实现了该接口，且同时还提供了缓冲和一些文本I/O的帮助函数的对象。
// 参考:https://blog.csdn.net/preyta/article/details/80655736
func main()  {
	// An artificial input source.
	const input = "1234 5678 1234567901234567890"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Create a custom split function by wrapping the existing ScanWords function.
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		//ScanBytes，按照byte进行拆分
		//ScanLines，按照行(“\n”)进行拆分
		//ScanRunes，按照utf-8字符进行拆分
		//ScanWords，按照单词(" ")进行拆分
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}
	// Set the split function for the scanning operation.
	scanner.Split(split)
	// Validate the input
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}

