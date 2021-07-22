package main

import (
	"fmt"

)

func AddFunc() {
	var s = []int{1, 2}
	if len(s) >= 3 {
		//	fmt.Println(s[:3])
		fmt.Println(s[3:])
	}
}
func main() {
	AddFunc()
	assert.Equal(&testing.T{}, err, nil)
}
