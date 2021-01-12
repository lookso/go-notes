package main

import (
	"fmt"
	"strings"
)

// 最长公共前缀
func main() {
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
