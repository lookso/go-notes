package main

import "fmt"

func main() {
	s := []int{1,}
	s = append(s, []int{2, 3, 4}...)
	a := s
	a[0] = 2
	fmt.Println(s) //  [2 2 3 4],a,s 引用自的是同一个数组,
	var b = make([]int, 10)
	copy(b, a) //  使用copy开辟一个全新的数组
	b[0] = 100
	fmt.Println(s) // [2 2 3 4]

	m:=new(struct{name string})
	m.name="jack"

	p := new([]int) //p == nil; with len and cap 0
	fmt.Println(p)
	v := make([]int, 10, 50) // v is initialed with len 10, cap 50
	fmt.Println(v)

	arr:=new([20]int)
	fmt.Println(cap(arr),len(arr))
}
