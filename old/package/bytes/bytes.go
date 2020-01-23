package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	state := bytes.EqualFold([]byte("abc"), []byte("abC"))
	fmt.Println(state)
	fmt.Println(bytes.HasSuffix([]byte("abc"), []byte("c")))
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("Go")))

	equalStatus := bytes.Equal([]byte("nb"), []byte("nb")) // 注意区别大小写
	if equalStatus {
		fmt.Println("equal", equalStatus)
	}
	if bytes.EqualFold([]byte("Nba"),[]byte("nba")) {
		fmt.Println("==")
	}
	if strings.EqualFold("Nba","nba"){
		fmt.Println("==")
	}
}
