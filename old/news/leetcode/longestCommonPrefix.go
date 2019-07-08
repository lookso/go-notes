package main

import (
	"fmt"
)

//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
//
//示例 1:
//
//输入: ["flower","flow","flight"]
//输出: "fl"
//示例 2:
//
//输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//说明:
//
//所有输入只包含小写字母 a-z 。

func main(){
	// fmt.Println(strings.Index("chicken", "ken")) // 不存在返回-1
	str:=longestCommonPrefix([]string{"flower","flow","flight"})

	fmt.Println("str:",str)
}
func longestCommonPrefix(strs []string) string{
	//if len(strs)==0 {
	//	return ""
	//}
	//prefix:=strs[0]
	//for i:=1;i<len(strs);i++{
	//	fmt.Println(strs[i])
	//	for strings.Index(prefix,strs[i])!=-1 {
	//		prefix = prefix[:len(prefix)-1]
	//		if prefix==""{
	//			return ""
	//		}
	//	}
	//}
	//return prefix

	s := ""
	n := 0      // 字符游标，用于判断所有字符串在游标位置的字符是否都相同，游标递增
	for {
		i := 0
		for ; i<len(strs); i++ {    // 遍历字符串数组
			if n >= len(strs[i]) {  // 如果当前游标超出某个字符串的长度，退出遍历
				break
			}
			if strs[i][n] != strs[0][n] {   // 如果当前位置的字符有不相同的，退出遍历
				break
			}
		}
		if len(strs)==0 || i<len(strs) {    // 判断是否是中途退出，如果是，说明已经找到最长公共前缀，退出函数
			return s
		} else {                            // 反之，游标递增，继续看下一个字符
			s = strs[0][:n+1]
			fmt.Println(s)
			n++
		}
	}
}