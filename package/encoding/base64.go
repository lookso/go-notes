package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"reflect"
)

func main() {
	// base64.StdEncoding 和base64.URLEncoding 的区别
	// 由于标准的Base64编码后可能出现字符+和/,在URL中就不能直接作为参数,
	// 所以又有一种"url safe"的base64编码,其实就是把字符 + 和 / 分别变成 - 和 _
	input := []byte("\x00\x00\x3e\x00\x00\x3f\x00")
	encodeToString := base64.StdEncoding.EncodeToString(input) // 编码
	fmt.Println(encodeToString)
	data, err := base64.StdEncoding.DecodeString(encodeToString) // 解码
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%q\n", data)

	// 如果要用在url中，需要使用URLEncoding
	uEnc := base64.URLEncoding.EncodeToString([]byte(input))
	fmt.Println(uEnc)

	uDec, err := base64.URLEncoding.DecodeString(uEnc)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(uDec))
	fmt.Println(reflect.TypeOf(uDec))

	//AAA+AAA/AA==
	//	"\x00\x00>\x00\x00?\x00"
	//AAA-AAA_AA==
	//>?
}
