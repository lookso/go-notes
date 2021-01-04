/*
@Time : 2019/1/8 9:49 AM
@Author : Tenlu
@File : input
@Software: GoLand
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

//  echo "haha"|go run input.go
//  haha
