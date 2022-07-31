package main

import (
	"fmt"
	"strings"
)

// 最长公共前缀
func main() {
	fmt.Println( strings.Index("abc", "abc"))
	var strs = []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	//fmt.Println(strings.Index("flower", "fl")) //0
	prefix := strs[0]
	for _, k := range strs {
		for strings.Index(k, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

func longCommonPrefix2(strs []string) string {
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

