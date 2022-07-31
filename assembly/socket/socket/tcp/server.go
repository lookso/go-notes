/*
@Time : 2019/3/19 11:51 AM 
@Author : Tenlu
@File : socket
@Software: GoLand
*/
package main

import (
	"net"
	"fmt"
	"os"
	"time"
)

//错误处理函数
func checkServerErr(err error, extra string) bool {
	if err != nil {
		formatStr := " Err : %s\n"
		if extra != "" {
			formatStr = extra + formatStr
		}

		fmt.Fprintf(os.Stderr, formatStr, err.Error())
		return true
	}

	return false
}

//连接处理函数
func svrConnHandler(conn net.Conn) {
	fmt.Println("Client connect success :", conn.RemoteAddr().String())
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	request := make([]byte, 128)
	defer conn.Close()
	for {
		readLen, err := conn.Read(request)
		if checkServerErr(err, "Read") {
			break
		}
		//socket被关闭了
		if readLen == 0 {
			fmt.Println("Client connection close!")
			break
		} else {
			//输出接收到的信息
			fmt.Println(string(request[:readLen]))
			time.Sleep(time.Second)
			//发送
			conn.Write([]byte("World !"))
		}
		request = make([]byte, 128)
	}
}

func main() {
	//解析地址
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:6666")
	if checkServerErr(err, "ResolveTCPAddr") {
		return
	}
	//设置监听地址
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if checkServerErr(err, "ListenTCP") {
		return
	}

	for {
		//监听
		fmt.Println("Start wait for client.")
		conn, err := listener.Accept()
		if checkServerErr(err, "Accept") {
			continue
		}

		//消息处理函数
		go svrConnHandler(conn)
	}
}
