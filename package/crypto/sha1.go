/*
@Time : 2019/5/25 10:56 PM 
@Author : Tenlu
@File : shar1
@Software: GoLand
*/
package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func main()  {
	hash := sha1.New()
	hash.Write([]byte("helloworld")) // 需要加密的字符串为 123456
	cipherStr := hash.Sum(nil)
	fmt.Printf("%s\n", hex.EncodeToString(cipherStr)) // 输出加密结果
}

