/*
@Time : 2019/5/27 11:23 AM 
@Author : Tenlu
@File : strackclose
@Software: GoLand
*/
package main

import (
	"strings"
	"fmt"
)

func main(){
	var str []string
	str = []string{"jack","toms","jerry"}
	str = append(str[:len(str)-1])
	fmt.Println("str:",str)

	status:=straceClose("(()))")
	fmt.Println(status)
}
func straceClose(str string) bool {
	if str=="" {
		return false
	}
	// 判断子字符串或字符在父字符串中出现的位置（索引) -1 表示不存在
	if strings.Index(str,"(")<0 && strings.Index(str,")")<0{
		return false
	}
	var strack []string
	for _,v:=range str{

		if string(v) == "(" {
			strack=append(strack, "(")
		}
		if string(v) == ")" {
			if len(strack)==0{
				return false
			}
			strack=append(strack[:len(strack)-1])
		}
	}
	if len(strack)!=0{
		return false
	}
	return true
}

