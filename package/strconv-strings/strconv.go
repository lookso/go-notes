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
	fmt.Println(strconv.IntSize)
	var id = 10
	// int 转int64
	numInt64 := int64(id)
	fmt.Println("int to int64", numInt64)
	// int64转 int
	Int64 := int(numInt64)
	fmt.Println("int64 to int", Int64)

	// int 转string
	numString := strconv.Itoa(id)
	fmt.Println("int to string", numString)
	// string 转int
	Int, _ := strconv.Atoi(numString)
	fmt.Println("string to int", Int)

	// int64转string
	// 返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母'a'到'z'表示大于10的数字。
	numString64 := strconv.FormatInt(64, 10)
	fmt.Println("int64 to string", numString64)
	//string 转int64
	idInt64, _ := strconv.ParseInt(numString, 10, 64)
	fmt.Println("string to int64", idInt64)

	china := map[string]interface{}{}
	china["bj"] = "10.01"

	fmt.Println(reflect.TypeOf(china["bj"]))
	if reflect.TypeOf(china["bj"]).String() == "string" {
		fmt.Println("string")
	}
	if value, ok := china["bj"].(string); ok {
		fmt.Println(value)
	}
}
