/*
@Time : 2019/5/18 10:23 AM 
@Author : Tenlu
@File : channel
@Software: GoLand
*/
package main

import "fmt"

func main()  {
	var tag =make(chan bool,10)

	go func() {
		fmt.Printf("channel cap:%d\n",cap(tag))
		tag<-true
	}()
	<-tag
}
