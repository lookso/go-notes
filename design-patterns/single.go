/*
@Time : 2019-07-24 17:24
@Author : Tenlu
@File : single
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	Name string
}

var ins *singleton

var once sync.Once

func GetIns() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}

func main() {
	singleA := GetIns()
	singleA.Name = "toms"
	singleB := GetIns()
	singleB.Name = "jack"
	fmt.Println(singleA.Name)  //
	fmt.Println(singleB.Name) // 我们申请了两次实例.在改变一个第二个实例的字段之后，第一个也随之改变了
}
