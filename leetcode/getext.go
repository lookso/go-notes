/*
@Time : 2019-10-28 18:52 
@Author : Tenlu
@File : getext
@Software: GoLand
*/
package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	var Url = "https://www.cnblogs.com/benlightning/articles/4441027.html"
	ext := getext(Url)
	fmt.Println(ext)
}
func getext(Url string) string {
	if urlInfo, err := url.Parse(Url); err == nil {
		pathStr := urlInfo.Path
		index := strings.Index(pathStr, ".")
		if index > 0 {
			return pathStr[index+1:]
		}
	}
	return ""
}
