package main

import (
	"bytes"
	"fmt"
)

func main()  {
	fmt.Println(bytes.Compare([]byte{100}, []byte{1, 2, 3, 1, 4}))
}
