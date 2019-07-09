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
	Author string `json:"author,omitempty"`
	Pages  int    `json:"pages,omitempty"`
}
func main() {
	var books Books
	fmt.Println(books)
}
