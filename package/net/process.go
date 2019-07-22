/*
@Time : 2019-07-21 10:44 
@Author : Tenlu
@File : process
@Software: GoLand
*/
package main

import (
	"fmt"
	"os"
)
func main(){
	fmt.Println(os.Getpid())
	os.Pipe()
}
