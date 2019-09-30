package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	fmt.Println(md5Func("213456"))
}
func md5Func(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
