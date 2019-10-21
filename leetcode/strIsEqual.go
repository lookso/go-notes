/*
@Time : 2019-10-20 10:38 
@Author : Tenlu
@File : strIsEqual
@Software: GoLand
*/

package main

import (
	"fmt"
	"strings"
)

func isUniqueStr(s string) bool {
	
	if len([]rune(s)) > 3000 {
		return false
	}
	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}

func main() {
	s1 := "aaaa"
	fmt.Println(isUniqueStr(s1))

	s2 := "BarackObama"
	fmt.Println(isUniqueStr(s2))
}
