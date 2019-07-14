/*
@Time : 2019-07-09 22:04 
@Author : Tenlu
@File : struct
@Software: GoLand
*/
package main

import (
	"fmt"
	"encoding/json"
)

type Books struct {
	Title  string `json:"title"`
	Author string `json:"author,omitempty"` // 为空则不输出
	Pages  int    `json:"pages,omitempty"`
}

func main() {
	b := new(Books)
	b.Title = "baidu"
	fmt.Println(b)

	var books Books
	fmt.Println(books)
	books.Title = "time comming"
	books.Author = "lk"
	books.Pages = 160
	fmt.Println(books)
	bookJson,err:=json.Marshal(&books)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("struct to json:",string(bookJson))
}
