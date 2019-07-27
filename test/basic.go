/*
@Time : 2019-07-27 21:50 
@Author : Tenlu
@File : basic
@Software: GoLand
*/

/*
go test -v  basic_test.go basic.go


*/
package main

import (
	"errors"
	"fmt"
	"sync"
)

func main()  {
	var group =new(sync.WaitGroup)
	var mutux =new(sync.Mutex)
	mutux.Lock()
	group.Add(2)
	go func() {
		for i:=1;i<6 ;i++  {
			fmt.Println(i)
		}
		mutux.Unlock()
		mutux.Lock()
		defer group.Done()
	}()
	go func() {
		for i:=7;i<14 ;i++  {
			fmt.Println(i)
		}
		mutux.Unlock()
		defer group.Done()

	}()
	group.Wait()


}
func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}

	return a / b, nil
}