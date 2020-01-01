package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(int64(math.Pow(10,8))) // 10的8次方
	fmt.Println(math.Abs(-10))
	fmt.Println(math.Ceil(4.6)) // 返回不小于x的最大整数（的浮点值），特例如下
	fmt.Println(math.Floor(5.13)) // 返回不大于x的最大整数（的浮点值），特例如下
	fmt.Println(math.Round(8.51)) // 四舍五入
	fmt.Println(math.Max(10,30))
}
