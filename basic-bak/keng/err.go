package main

import "fmt"

type DetailError struct {
	code    int
	message string
}

// 这个方法没用到为啥会被调用？var err error
func (e DetailError) Error() string {
	return fmt.Sprintf("error (%d and %s)", e.code, e.message)
}
func handle(x int) *DetailError {
	if x != 1 {
		fmt.Println(&DetailError{code: 10001, message: "error msg"})
	}
	return nil
}
func String()  {
	fmt.Println("nb-string")
}
func main() {
	var err error
	err = handle(0)
	if err != nil {
		fmt.Println("the error 1:", err)
	}
	err = handle(1)
	if err != nil {
		fmt.Println("the error 2:", err)
	}
}
