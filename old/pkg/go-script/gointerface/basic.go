/*
@Time : 2019/3/3 9:28 PM 
@Author : Tenlu
@File : basic
@Software: GoLand
*/
package main

import (
	"fmt"
	"os"
)

type people interface {
	run() (peopleChan chan<- *man)
	eat() string
}
type man struct {
	name string
	age  int
}

//
//var m = man{name: "jack", age: 12}
var peopleChan = make(chan string)

func (r *man) run() (peopleChan chan<- *man) {
	peopleChan <- &man{name: "toms", age: 10}
	return peopleChan
}

func (r *man) eat() string {
	return r.name
}
func main() {
	var pArr = []people{
		&man{name: "jack", age: 10},
	}
	//for _, p := range pArr {
	//	//run := <-p.run()
	//	run:=p.run()
	//	fmt.Printf("%T",run)
	//	os.Exit(12)
	//	//fmt.Printf("hello,my name is %s,my age is %\n", p.eat(), run)
	//
	//}

	close(peopleChan)
}
