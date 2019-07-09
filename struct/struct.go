/*
@Time : 2019-07-09 22:04 
@Author : Tenlu
@File : struct
@Software: GoLand
*/
package main

import "fmt"

type Books struct {
	Title  string `json:"title"`
	Author string `json:"author,omitempty"` // 为空则不输出
	Pages  int    `json:"pages,omitempty"`
}

func main() {
	var books Books
	fmt.Println(books)
	books.Title = "time comming"
	books.Author = "lk"
	fmt.Println(books)
}
