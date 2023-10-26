package main

import (
	"fmt"
	"github.com/pjebs/optimus-go"
)

func main() {
	// 创建 Optimus 实例
	o := optimus.New(876339253, 1776316957, 835908656)

	// 假设我们有一个 Optimus ID
	optimusID := uint64(1103647398)

	// 解码 Optimus ID
	originalID := o.Decode(optimusID)

	fmt.Printf("Original ID: %d", originalID)
}
