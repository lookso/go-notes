package main

import (
	"errors"
	"fmt"

	"github.com/spf13/cast"
	"golang.org/x/sync/errgroup"
)

func main() {
	var arr = []int{1, 2, 3}
	g := errgroup.Group{}
	for _, v := range arr {
		nv := v // 创建一个新的变量来保存 v 的值
		//	fmt.Printf("v:%d\n", v)
		g.Go(func() error {
			// 在这里使用 v
			t, err := myt(nv)
			fmt.Println(t)
			return err
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
func myt(v int) (string, error) {
	fmt.Println(v)
	t := cast.ToString(v) + "dnb"
	return t, errors.New("fuck")
}
