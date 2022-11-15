package main

import (
	"fmt"
	"github.com/edsrzf/mmap-go"
	"os"
)

func main() {
	f, _ := os.OpenFile("./log/main.go", os.O_RDWR, 0644)
	defer f.Close()

	mmap, _ := mmap.Map(f, mmap.RDWR, 0)
	defer mmap.Unmap()
	fmt.Println(string(mmap))
	mmap.Flush()
}
