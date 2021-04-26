package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "中国 牛逼"
	if Sp(str){
		fmt.Println("ok")
	}else{
		fmt.Println("false")
	}
}

func Sp(str string) bool {
	
	for _, v := range str {
		if  unicode.IsSpace(v) {
			return true
		}
	}
	return false
}
