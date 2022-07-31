/*
@Time : 2019/3/11 11:37 AM
@Author : Tenlu
@File : bitmap
@Software: GoLand
*/

//http://blog.studygolang.com/2014/09/bitmap_multi_language/

package main

import (
	"fmt"
)

// 问题1:假如给你20亿个非负数的int型整数,然后再给你一个非负数的int型整数t,让你判断t是否存在于这20亿数中,你会怎么做呢？
// 问题2:位图排序

var nums = [100]int{11, 2, 3, 20, 10, 22, 42, 67, 8, 23, 2, 30}

func main() {
	bitMap(1, nums)
}

func bitMap(n int, nums [100]int) {
	// 时间复杂度 O(n)
	/*
	想一下，这样的话，时间复杂度是O(n)，所需要的内存空间
	4byte * 20亿，一共需要80亿个字节，
	大概需要8GB的内存空间，显然有些计算机的内存一次是加载不了这么这么多的数据的。
	*/
	//for _, v := range nums {
	//	if n == v {
	//		fmt.Printf("the num %d exists\n", n)
	//		break
	//	} else {
	//		fmt.Printf("the num %d not exists\n", n)
	//		break
	//	}
	//}
	// 初步优化:时间复杂度O(1)
	if nums[n] != 0 {
		fmt.Printf("the num %d exists\n", n)
	}
}
