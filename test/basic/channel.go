package main

import "fmt"

func main()  {
	var ch =make(chan int)
	close(ch)
	fmt.Println(<-ch)

	//var chMain=make(chan bool,0)
	//
}
