package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func main()  {
	h := sha1.New()
	io.WriteString(h, "this pwd 123456")
	fmt.Printf("%x\n", h.Sum(nil))
	//SHA1 值经常以 16 进制输出，使用%x 来将散列结果格式化为 16 进制字符串。
	data := []byte("This page intentionally left blank.")
	fmt.Printf("SHA1 : %x\n",sha1.Sum(data))
}
