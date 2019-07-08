/*
@Time : 2019/5/15 11:51 AM 
@Author : Tenlu
@File : list
@Software: GoLand
*/
package main

import (
	"container/list"
	"fmt"
)
// 对于很多数据来讲:频繁的插入和删除用list,频繁的遍历查询选slice
// 当程序要求slice的容量超大并且需要频繁的更改slice的内容时，就不应该用slice，改用list更合适
func main()  {
	var list = list.New()
	list.PushFront("12")
	fmt.Printf("list's len:%d\n",list.Len())
	fmt.Printf("list's value:%v\n",list.Front().Value)
}

