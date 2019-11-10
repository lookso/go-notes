// +build windows darwin

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	l1, err := net.Listen("tcp", os.Args[2])
	if err != nil {
		panic(fmt.Errorf("net.Listen(\"tcp\", %v) error(%v)", os.Args[2], err))
	}

	list := append([]net.Listener{}, l1)

	fmt.Println(list[0].Addr())
	
	//// 加载初始化值
	//Init(os.Args[1], list)
	//defer Quit()
	//
	//c := make(chan os.Signal, 1)
	//signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	//for {
	//	s := <-c
	//	switch s {
	//	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
	//		return
	//	case syscall.SIGHUP:
	//		// TODO reload
	//	default:
	//		return
	//	}
	//}

}
