package main

// https://leetcode-cn.com/problems/string-to-integer-atoi/solution/go-chang-gui-bian-li-by-ba-xiang/
import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("-123"))
}
func myAtoi(str string) int {
	res := 0
	tag := 1
	for k, s := range str {
		if s >= '0' && s <= '9' {
			res = res*10 + int((s - '0'))
			if res*tag > math.MaxInt32 {
				return math.MaxInt32
			}
			if res*tag < math.MinInt32 {
				return math.MinInt32
			}
		} else if s == '-' {
			if k != 0 || len(str) == 1 {
				return -1
			}
			tag = -1
		} else {
			return -1
		}
	}
	return res * tag

}
