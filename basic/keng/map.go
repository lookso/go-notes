package main

import (
	"fmt"
)

type Student struct {
	Age int
}

func main() {
	//var kv  = make(map[string]Student,1000)
	//kv := map[string]Student{"menglu": {Age: 21}}
	//kv["menglu"].Age = 22 // 报错

	//s := []Student{{Age: 21}}
	//s[0].Age = 22
	//fmt.Println(kv, s)
	//x = y 这种赋值的方式，你必须知道 x的地址，然后才能把值 y 赋给 x。
	//但 go 中的 map 的 value 本身是不可寻址的，因为 map 的扩容的时候，可能要做 key/val pair迁移
	//value 本身地址是会改变的
	// 问题描述:https://gocn.vip/question/1714?uid=focus
	// map的v不可寻址是因为map的扩容的过程中会进行一个拷贝 （capacity很重要，减少拷贝），
	// 但是map的v是一个引用类型是可以更改的吧。我的理解。

	mapT := map[string]map[string]string{"name": {"num1":"aaa"}}
	mapT["name"]["num1"]="bbb"
	fmt.Println(mapT)


	i := 2
	if i > 1 {
		i, err := doTest(i, 2)
		fmt.Println(&i)
		if err != nil {
			panic(err)
		}
		fmt.Println(&i)
	}
	fmt.Println(&i)
	fmt.Println("------j-----")
	j:=2
	if j>1{
		j:=1
		fmt.Println(j)
	}
	fmt.Println(j)
}
func doTest(x,y int)(int,error)  {
	return x/y,nil
}
