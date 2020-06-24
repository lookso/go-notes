package main
import (
	"fmt"
	"github.com/yuin/gopher-lua"
)

func main()  {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile("hello.lua"); err != nil {
		panic(err)
	}
	lv := L.Get(-1)
	fmt.Println(lv)
}
