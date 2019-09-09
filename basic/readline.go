/*
@Time : 2019-09-07 22:54 
@Author : Tenlu
@File : readline
@Software: GoLand
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	fi, err := os.Open("/Users/lukun/data/apps/works/go-apps/src/source-code/echodemo/main.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		fmt.Println(string(a))
	}
}
