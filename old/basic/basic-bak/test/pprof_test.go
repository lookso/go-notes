/*
@Time : 2019-07-23 16:53
@Author : Tenlu
@File : pprof
@Software: GoLand
*/
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println(123)
	}()
	http.ListenAndServe("localhost:6060", nil)
}
