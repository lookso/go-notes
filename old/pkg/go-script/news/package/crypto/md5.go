/*
@Time : 2019/5/25 10:56 PM 
@Author : Tenlu
@File : md5
@Software: GoLand
*/
package main

import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
)

func main()  {
	hash := md5.New()
	hash.Write([]byte("123456")) // 需要加密的字符串为 123456
	cipherStr := hash.Sum(nil)
	fmt.Println(cipherStr)
	// 字节转16进制
	fmt.Printf("%s\n", hex.EncodeToString(cipherStr)) // 输出加密结果
}

