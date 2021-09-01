package main

import (
	"fmt"
	"hash/fnv"
)

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// 将字符串hash为数字
func main() {
	fmt.Println(hash("HelloWorld")) // 926844193
	fmt.Println(hash("Hello"))      // 4116459851

	count := 1000
	fmt.Println(hash("https://baidu.com"), hash("https://baidu.com")%uint32(count))
	fmt.Println(hash("https://baidu.com/php/7.8"), hash("https://baidu.com/php/7.8")%uint32(count))
	fmt.Println(hash("https://360.com"), hash("https://360.com")%uint32(count))
	fmt.Println(hash("https://yahoo.com"), hash("https://yahoo.com")%uint32(count))
	fmt.Println(hash("https://google.com"), hash("https://google.com")%uint32(count))
	fmt.Println(hash("https://bing.com/go"), hash("https://bing.com/go")%uint32(count))
	fmt.Println(hash("https://yahoo.com"), hash("https://yahoo.com")%uint32(count))
}
