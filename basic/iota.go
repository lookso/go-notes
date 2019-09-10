package main

import "fmt"

// 每遇到一个const关键字,iota就会重置为0,const中每新增一行常量声明将使iota计数一次
const (
	Num0 = iota + 100
	Num1
	Num55 = iota
	Num66
	_ // 跳过值
	Num3 = "zz"
	Num4
	Num5 = iota
	Num6
)

// 定义在一行的情况,iota 在下一行增长，而不是立即取得它的引用。
const (
	Apple, Banana = iota + 100, iota + 200
	Cherimoya, Durian
	Elderberry, Fig
	Num = iota
)

// 中间插队
const (
	i = iota
	j = 3.14
	k = iota
	l
)
const (
	Aa, Bb = iota + 100, iota + 200
	// Cc 会报错,数量得保持一致
	Cc, Dd
)

func main() {
	fmt.Println(Num0, Num1, Num55, Num66, Num3, Num4, Num5, Num6) // 100 101 2 3 zz zz 7 8
	fmt.Println("--------")
	fmt.Println(Apple, Banana, Cherimoya, Durian, Elderberry, Num) // 100 200 101 201 102 3
	fmt.Println("--------")
	fmt.Println(i, j, k, l) // 0 3.14 2 3
	fmt.Println("--------")
	fmt.Println(Aa, Bb, Cc) // 0 3.14 2 3

}
