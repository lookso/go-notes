/*
@Time : 2019-11-02 09:55 
@Author : Tenlu
@File : utf8
@Software: GoLand
*/
package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func main() {
	fmt.Println(utf8.RuneCount([]byte("nba")))
	fmt.Println(utf8.RuneCountInString("nba"))
	// 报告s是否包含完整且合法的utf-8编码序列。
	valid := "Hello, 世界"
	invalid := string([]byte{0xff, 0xfe})
	fmt.Println(utf8.ValidString(valid))   // true
	fmt.Println(utf8.ValidString(invalid)) //false

	// EncodeRune将r的utf-8编码序列写入p（p必须有足够的长度），并返回写入的字节数。
	r := '世'
	buf := make([]byte, 3)
	n := utf8.EncodeRune(buf, r)
	fmt.Println(buf) // [228 184 150]
	fmt.Println(n)

	r, size := utf8.DecodeRuneInString("nba")
	fmt.Println(r, size)

	var buffer bytes.Buffer
	for i := 0; i < 10; i++ {
		buffer.WriteString("a")
	}
	buffer.WriteString("nb")
	buffer.WriteString("china")
	fmt.Println(buffer.String())

	//p:=unsafe.Pointer(&user)
	//fmt.Println(p)

	ndb := NewDb()
	ndb.Conn()
}

type Db struct {
	Name string
}

func (db *Db) Conn() {
	fmt.Println(&db)
	fmt.Println(unsafe.Pointer(&db))
}
func NewDb() *Db {
	return &Db{}
}
