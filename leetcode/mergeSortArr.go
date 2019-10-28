/*
@Time : 2019-10-28 16:08 
@Author : Tenlu
@File : mergesortarr
@Software: GoLand
*/
package main

import "fmt"

// 合并2个有序数组
func main() {
	var a = []int{1, 2, 3, 4, 5}
	var b = []int{2, 3, 4, 5, 6}

	c := sortArr(a, b)
	fmt.Println("c:", c)
	//
	//|| 1 2
	//-- 2 2
	//|| 2 3
	//-- 3 3
	//|| 3 4
	//-- 4 4
	//|| 4 5
	//-- 5 5
	//|| 5 6
	//aindex 5
	//bindex 4
	//cindex 9
	//[1 2 2 3 3 4 4 5 5 0]
	//c: [1 2 2 3 3 4 4 5 5 6]

}

func sortArr(a, b []int) []int {

	//计算数组长度
	alen := len(a)
	blen := len(b)
	clen := alen + blen

	//创建结果集数组
	c := make([]int, clen)

	//默认数组从0开始
	aIndex := 0
	bIndex := 0
	cIndex := 0

	// 开始遍历数组
	for aIndex < alen && bIndex < blen {

		if a[aIndex] < b[bIndex] {
			fmt.Println("||", a[aIndex], b[bIndex])
			c[cIndex] = a[aIndex]
			cIndex++
			aIndex++
		} else {
			fmt.Println("--", a[aIndex], b[bIndex])
			c[cIndex] = b[bIndex]
			cIndex++
			bIndex++
		}
	}

	fmt.Println("aindex", aIndex)
	fmt.Println("bindex", bIndex)
	fmt.Println("cindex", cIndex)
	fmt.Println(c)
	// 当其中一个数组遍历完,另一个数组还没遍历完应.当将另一个数组遍历完
	for aIndex < alen {
		c[cIndex] = a[aIndex]
		cIndex++
		aIndex++
	}
	for bIndex < blen {
		c[cIndex] = b[bIndex]
		cIndex++
		bIndex++
	}
	//返回数组
	return c
}
