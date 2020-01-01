/*
@Time : 2019-08-09 22:12 
@Author : Tenlu
@File : defer
@Software: GoLand
*/
package main

import "fmt"

/*
1.return对返回变量赋值，如果是匿名返回值就先声明再赋值；
2.执行defer函数；
3.return携带返回值返回。

*/

func deferFunc() int {
	index := 0

	fc := func() {

		fmt.Println(index, "匿名函数1")
		index++

		defer func() {
			fmt.Println(index, "匿名函数1-1")
			index++
		}()
	}

	defer func() {
		fmt.Println(index, "匿名函数2")
		index++
	}()

	defer fc()

	return func() int {
		fmt.Println(index, "匿名函数3")
		index++
		return index
	}()
}

func main() {
	fmt.Println(deferFunc())
	fmt.Println("v:", deferTest())
	//deferTest()
}
func deferTest() int {
	x := 1
	defer func(x int) {
		x = 100
		fmt.Println("defer x:", x)
	}(x)
	return func(x int) int {
		x = 200
		return x
	}(x)

}

func tt() int {
	return 123
}
