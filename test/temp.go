package main

import (
	"fmt"
	"github.com/spf13/cast"
	"math"
	"strconv"
)

// 一组数字,按照每5个元素长度起一个协程,不足5个的按5个算,有序打印出所有元素
func main() {
	var arr = []int{20, 21, 40, 12, 435, 4, 32, 17, 34, 765, 678, 23, 29}
	var newArr = make([]int, 0)
	var m = map[int]int{4: 1}
	//var temp=4
	for k, v := range arr {
		newArr = append(newArr, v)
		if _, ok := m[v]; ok {
			newArr[0], newArr[k] = v, newArr[0]
		}
	}
	// 401  3 1203(500+500+203)

	// 200  3  600
	// 30 4 120

	lf, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(1203)/float64(500)), 64)
	l := math.Ceil(lf)
	ll := cast.ToInt(l)
	// 0-166
	// 166-332
	//332-(401-332)
	ll = 3
	for i := 1; i <= ll; i++ {
		if ll > 1 {
			s := 500 / 3
			if i == ll {
				fmt.Println((i-1)*s,":",401)
			} else {
				fmt.Println((i-1) * s,":",i * s)
			}
		}
	}

	fmt.Println(newArr)

}
