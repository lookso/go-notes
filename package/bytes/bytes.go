package main

import (
	"bytes"
	"fmt"
)

func main() {
	state := bytes.EqualFold([]byte("abc"), []byte("abC"))
	fmt.Println(state)
	fmt.Println(bytes.HasSuffix([]byte("abc"), []byte("c")))
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("Go")))
}
