/*
@Time : 2019/5/25 10:56 PM 
@Author : Tenlu
@File : shar1
@Software: GoLand
*/
package main

import (
	"fmt"
	"encoding/hex"
	"crypto/sha1"
)

func main()  {
	hash := sha1.New()
	hash.Write([]byte("123456")) // 需要加密的字符串为 123456
	cipherStr := hash.Sum(nil)
	fmt.Printf("%s\n", hex.EncodeToString(cipherStr)) // 输出加密结果
}

