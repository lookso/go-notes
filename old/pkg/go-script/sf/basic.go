/*
@Time : 2019/3/28 6:56 PM 
@Author : Tenlu
@File : basic
@Software: GoLand
*/
package main

import (
	"errors"
	"fmt"
	"strings"
)

func main()  {

	status,err:=strackClose(")()")
	if err != nil {
		fmt.Printf("err:%s\n", err)
		return
	}
	println(status)

	fmt.Println(twoSum([]int{5,2,1,6,7,8,9,10,5,4},10))

	fmt.Println(test())
}

// 1.写个函数,判断下面扩号是否闭合,左右对称即为闭合:((())),)(()),(())))
func strackClose(str string) (b bool,err  error){
	if str=="" {
		err = errors.New("str empty")
		return
	}
	// 判断子字符串或字符在父字符串中出现的位置（索引) -1 表示不存在
	if strings.Index(str,"(")<0 && strings.Index(str,")")<0{
		err = errors.New("params ( or ) not exists")
		return
	}
	var strack []string
	// golang中字符类型的实际数据类型为uint32
	// golang中string底层是通过byte数组实现的。
	// 中文字符在unicode下占2个字节,在utf-8编码下占3个字节,而golang默认编码正好是utf-8。

	for _,v:=range str{
		if string(v) == "(" {
			strack=append(strack, "(")
		}
		if string(v) == ")" {
			if len(strack)==0{
				return false,nil
			}
			strack=append(strack[:len(strack)-1])
		}
	}
	if len(strack)!=0{
		return false,nil
	}
	return true,nil
}

// 2.从数组中找出和为目标值的两个元素,[]int{5,2,1,6,7,8,9,10,5,4}
func twoSum(arr []int,sum int) []int {
	if len(arr)<2{
		return nil
	}
	var result []int
	for _,v:=range arr {
		minus := sum-v
		for _,vv:=range arr{
			if minus==vv {
				result = append(result, minus)
			}
		}
	}
	return result
}

func test() []int {
	println("---------\n")
	var a []int
	a=[]int{}
	for i:=0;i<10 ;i++  {
		a =append(a,i)
	}
	return a
}
