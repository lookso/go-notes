package main

import (
	"encoding/json"
	"fmt"
	"time"
)

//- 切片、映射和通道，使用make
//- 数组、结构体和所有的值类型，使用new

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a()
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}

// 结构体-方法
type TwoInts struct {
	a int
	b int
}

type Address struct {
	Type    string
	City    string
	Country string
}
type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}

func main() {
	defer func() {
		fmt.Printf("\a--run-end--\n")
	}()

	b()
	print("\n")
	callback(1, Add)

	// 匿名函数
	defer func() {
		sum := 0
		for i := 1; i <= 100; i++ {
			sum += i
		}
		fmt.Printf("高斯函数之和:%d\n", sum)
	}()

	start := time.Now()
	i := 0
	where := func() {
		i++
		fmt.Printf("hello,this is %d\n", i)
	}
	where()
	where()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)

	var arr = [...]string{"jack", 1: "polly", "robin"}
	for k, v := range arr {
		fmt.Printf("%d=>%s\n", k, v)
	}
	fmt.Printf("cap:%d,len:%d\n", cap(arr), len(arr))

	// 生成切片
	// make([]int, 50, 100)
	// new([100]int)[0:50]

	var numbers = make([]int, 4, 7)
	numbers = []int{1, 2, 3, 4, 8, 9}
	numbers = append(numbers, 4, 5, 6)

	//切片提供了计算容量的函数 cap()
	//可以测量切片最长可以达到多少:它等于切片的长度 + 数组除切片之外的长度。如果 s 是一个切片,
	//cap(s) 就是从 s[0] 到数组末尾的数组长度。切片的长度永远不会超过它的容量,
	//所以对于 切片 s 来说该不等式永远成立：0 <= len(s) <= cap(s)

	for j := 0; j < len(numbers); j++ {
		fmt.Printf("numbers:%d->%d\n", j, numbers[j])
	}
	// cap:12,len:9,type:[]int
	fmt.Printf("cap:%d,len:%d,type:%T\n", cap(numbers), len(numbers), numbers)

	// 结构体方法
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", two2.AddThem())

	//  处理json
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s\n", js)

	sum := 0
	for {
		sum ++
		if sum > 10 {
			break
		} else {
			fmt.Println(sum)
		}
	}
}
