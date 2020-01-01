/*
@Time : 2019/5/11 11:38 PM 
@Author : Tenlu
@File : basic
@Software: GoLand
*/
package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main(){
	// chan 类型的空值是 nil，声明后需要配合 make 后才能使用。
	var ch=make(chan interface{},10)
	ch<-10
	ch<-"jack"
	ch<-map[string]interface{}{"name":"jack","age":11}

	close(ch) // 此处不close通道,遍历时则报错
	fmt.Printf("%v\n",<- ch)
	for c:=range ch {
		fmt.Printf("for:%v\n",c)

		fmt.Sprintf("type:%T\n", c)
		if newValue,ok:=c.(map[string]interface{});ok{
			fmt.Printf("type:%s\n",reflect.TypeOf(c)) // 获取变量类型
			// map[string]interface {}
			if _,ok:=newValue["name"];ok{
				for k,v:=range newValue {
					fmt.Printf("%v->%v,\n",k,v)
				}
			}
		}
	}
	fmt.Printf("%v\n",<- ch) // nil

	var name="jack"
	fmt.Printf("%v\n",typeof(name))
	closeChan()

	go mm1()
	go mm2()
	for{
		runtime.GC()
	}
}
// 变量类型判断
func typeof(v interface{}) string {
	switch t := v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case float64:
		return "float64"
	default:
		_ = t
		return "unknown"
	}
}


func closeChan(){
	var c = make(chan int,10)
	go func() {
		for i:=0;i<10;i++  {
			c<-i
		}
		close(c)
	}()
	for{
		// 通道关闭
		if data,ok:=<-c;ok{
			fmt.Println(data)
		}else{
			break
		}
	}
}

var ch8 = make(chan int, 6)

func mm1() {
	for i := 0; i < 10; i++ {
		ch8 <- 8 * i
	}
	close(ch8)
}
func mm2() {
	for {
		for data:=range ch8{
			fmt.Print(data,"\t")
		}
	}
}
