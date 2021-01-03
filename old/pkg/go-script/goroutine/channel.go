/*
@Time : 2019/1/9 2:12 PM
@Author : Tenlu
@File : channel
@Software: GoLand
*/

//channel
//goroutine运行在相同的地址空间，因此访问共享内存必须做好同步。那么goroutine之间如何进行数据的通信呢，
//Go提供了一个很好的通信机制channel。channel可以与Unix shell 中的双向管道做类比：可以通过它发送或者接收值。
//这些值只能是特定的类型：channel类型。定义一个channel时，也需要定义发送到channel的值的类型。注意，必须使用make 创建channel：

//channel通过操作符<- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
//ch <- v    // 发送v到channel ch.
//v := <-ch  // 从ch中接收数据，并赋值给v

package main

import (
	"fmt"
	"time"
)

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}
func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	fmt.Printf("%v\n%v\n%v\n", a[:], a[2:], a[2:3])
	c := make(chan int)
	fmt.Printf("%v\n", a[:len(a)/2])
	fmt.Printf("%v\n", a[len(a)/2:])

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)

	fmt.Printf("---------------------------\n")
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)

}

//[7 2 8 -9 4 0]
//[8 -9 4 0]
//[8]
//[7 2 8]
//[-9 4 0]
//-5 17 12

//默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得Goroutines同步变的更加的简单，
//而不需要显式的lock。所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。其次，任何发送（ch<-5）将会被阻塞，直到数据被读出。无缓冲channel是在多个goroutine之间同步很棒的工具。

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getData(ch chan string) {
	var input string
	// time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%s\n", input)
	}
}

//Washington
//Tripoli
//London
//Beijing
//Tokio
