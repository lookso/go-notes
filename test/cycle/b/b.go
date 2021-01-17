package b

import (
	"fmt"
	"go-notes/test/cycle/c"
)

func Testb()  {
	fmt.Println("a")
	c.Testc()
}
