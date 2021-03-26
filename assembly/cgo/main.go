package main

/*
#include <stdint.h>

static int32_t add(int32_t a, int32_t b) {
    return a + b;
}
*/
import "C"
import (
	"fmt"
	"time"
)

func main() {
	var a, b = 1, 2
	start := time.Now().Nanosecond()
	for i := 1; i <= 6000; i++ {
		var c int32 = int32(C.add(C.int32_t(a+i), C.int32_t(b+i)))
		fmt.Println(i, c) // 3
	}
	end := time.Now().Nanosecond()
	runTime := end - start
	fmt.Println("run-time:", runTime)
}
