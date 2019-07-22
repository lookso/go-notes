/*
@Time : 2019-07-18 10:55 
@Author : Tenlu
@File : strconv
@Software: GoLand
*/
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var num int = 101
	// int转int64
	numInt64 := int64(num)
	// int转string
	numString := strconv.Itoa(num)
	// int64 转string
	int64ToString := strconv.FormatInt(numInt64, 10)
	// string转int64
	stringToInt64, _ := strconv.ParseInt(numString, 10, 64)
	fmt.Printf("int64tostring:%s,stringToInt64:%d\n", int64ToString, stringToInt64)
	fmt.Println("type:", reflect.TypeOf(stringToInt64))

	address := map[string]interface{}{}
	address["a"] = "beijing"
	
	fmt.Println(reflect.TypeOf(address["a"]))
	if reflect.TypeOf(address["a"]).String()=="string"{
		fmt.Println("string")
	}
	if value,ok:=address["a"].(string);ok{
		fmt.Println(value)
	}

	bool:=bool(true)
}
