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

type UserInfo struct {
	age int
}

func (u *UserInfo) setAge(age int) {
	u.age = age
}
func main() {
	var once sync.Once
	for i := 1; i < 10; i++ {
		once.Do(func() {
			fmt.Println(i)
		})
	}
	var groupOnce sync.Once
	var user = new(UserInfo)
	var group sync.WaitGroup
	group.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			// 只执行一次
			groupOnce.Do(func() {
				user.setAge(i)
				fmt.Println(user.age)
			})
			group.Done()
		}()
	}
	group.Wait()
}
