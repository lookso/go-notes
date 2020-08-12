package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 根节点
	rootCtx := context.Background()
	// 子节点
	ctx1 := context.WithValue(rootCtx, "name", "jack")
	// 子子节点,一层层向下传递
	ctx2, cancel := context.WithCancel(ctx1)
	defer cancel()
	name, ok := ctx2.Value("name").(string)
	if ok {
		fmt.Println(name)
	}


}
