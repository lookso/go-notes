/*
@Time : 2019-10-30 19:55
@Author : Tenlu
@File : strce
@Software: GoLand
*/
package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

//var strs = []string{"a", "b", "c"}
//var res = make([]string, 0)
//strLen := len(strs)
//for i := strLen - 1; i >= 0; i-- {
//res = append(res, strs[i])
//}
//fmt.Println(res)

func main() {

	status, err := strackClose("((()()")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("status:", status)
}

// 1.写个函数,判断下面扩号是否闭合,左右对称即为闭合:((())),)(()),(())))
func strackClose(str string) (b bool, err error) {

	if str == "" {
		err = errors.New("str empty")
		return
	}
	// 判断子字符串或字符在父字符串中出现的位置（索引) -1 表示不存在
	if strings.Index(str, "(") < 0 && strings.Index(str, ")") < 0 {
		err = errors.New("params ( or ) not exists")
		return
	}
	var strack []string
	// golang中字符类型的实际数据类型为uint32
	// golang中string底层是通过byte数组实现的。
	// 中文字符在unicode下占2个字节,在utf-8编码下占3个字节,而golang默认编码正好是utf-8。
	// )()()(
	var count = 0
	for _, v := range str {
		if string(v) == "(" {
			count += 1
			strack = append(strack, "(")
		}
		if string(v) == ")" {
			if len(strack) == 0 {
				return false, nil
			}
			count -= 1
			strack = strack[:len(strack)-1]
		}
	}
	fmt.Println("count", count)
	// ))()()((
	if len(strack) != 0 {
		return false, nil
	}
	return true, nil
}
