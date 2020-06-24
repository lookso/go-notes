package main

import (
	"fmt"
	"github.com/dengsgo/math-engine/engine"
	"math"
	"regexp"
)

func main() {
	// s := "1 + 2 * 6 / 4 + (456 - 8 * 9.2) - (2 + 4 ^ 5)"

	//s := "1******2-(1+(3-------1))///2"
	s := "1+1"
	fmt.Println("s", s)
	reg, _ := regexp.Compile("\\++")
	str := reg.ReplaceAllString(s, "+")

	reg2, _ := regexp.Compile("-+")
	str1 := reg2.ReplaceAllString(str, "-")
	fmt.Println("str1", str1)

	reg3, _ := regexp.Compile("\\*+")
	str2 := reg3.ReplaceAllString(str1, "*")
	fmt.Println("str2", str2)

	reg4, _ := regexp.Compile("/+")
	str3 := reg4.ReplaceAllString(str2, "/")
	fmt.Println("str2", str3)
	// call top level function
	r, err := engine.ParseAndExec(str3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s = %v\n", str3, r)

	fmt.Println("log100", math.Log10(100))

	var i int = 10

	fmt.Printf("Factorial of %d is %d\n", i, factorial(i))

	var res float64
	res=1
	fmt.Println("res",res)


	fmt.Println("Pow(2,3)",math.Pow(2,4))
}
// 阶乘
func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}
