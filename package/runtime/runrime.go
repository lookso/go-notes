package main

import (
	"fmt"
	"runtime"
)

func main() {
	defer func() {
		fmt.Printf("cpu-num:%+v,goroutine-num:%+v", runtime.NumCPU(), runtime.NumGoroutine())
	}()
	var done = make(chan int)
	for i := 0; i < 1; i++ {
		go func(done chan int, i int) {
			done <- i
		}(done, i)
		fmt.Println(<-done)
	}
	for i := 0 ; i< 4; i++ {
		test(i)
	}

}

func test(skip int) {
	call(skip)
}
func call(skip int) {
	// skip为0的时候表示当前所在的函数，即栈顶，1是从栈顶往下数第二个，以此类推，
	// line为执行了所在函数内的哪一行，
	// file为函数所在的文件名，
	// pc是所在函数的指针，
	pc,file,line,ok := runtime.Caller(skip)
	pcName := runtime.FuncForPC(pc).Name()  //获取函数名
	fmt.Println(fmt.Sprintf("res:%v   %s   %d   %t   %s",pc,file,line,ok,pcName))
	fmt.Println(runtime.Stack([]byte("123"),true))
}
