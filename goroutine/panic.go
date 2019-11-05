/*
@Time : 2019-10-30 22:42 
@Author : Tenlu
@File : panic
@Software: GoLand
*/
package main

import (
	"errors"
	"fmt"
)

func main() {
	var exit = make(chan bool)
	var errChan = make(chan error)
	go goPanic(exit, errChan)

	fmt.Println(<-errChan)
//	<-exit
}

func goPanic(exit chan bool, errChan chan error) {

	defer func() {
		if r := recover(); r != nil {
			errChan <- errors.New("some thing happened")
		}
	}()

	panic("num >10")
//	exit <- true
}
