package main

import "fmt"

func main() {
	myGoToFunc()
	myForFunc()
}

// goto语句
func myGoToFunc() {
	i := 0
Here: // 这行的第一个词，以冒号结束作为标签,标签名是大小写敏感的。
	println(i)
	i++
	if i > 100 {
		return
	}
	goto Here //跳转到Here去
}
func myForFunc() () {
	j := 0
	for {
		fmt.Println(j)
		j++
		if j > 100 {
			return
		}
	}
}
