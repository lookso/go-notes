
package main
import (
	"fmt"
	"os"
	"net"
	"sync"
)

// https://studygolang.com/articles/13907?fr=sidebar
var gLocker sync.Mutex    //全局锁
var gCondition *sync.Cond //全局条件变量

//错误处理函数
func checkErr(err error, extra string) bool {
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
func clientConnHandler(conn net.Conn) {
	gLocker.Lock()
	defer gLocker.Unlock()

	defer conn.Close()

	request := make([]byte, 128)
	for {
		readLen, err := conn.Read(request)
		if checkErr(err, "Read") {
			gCondition.Signal()
			break
		}

		//socket被关闭了
		if readLen == 0 {
			fmt.Println("Server connection close!")

			//条件变量同步通知
			gCondition.Signal()
			break
		} else {
			//输出接收到的信息
			fmt.Println(string(request[:readLen]))

			//发送
			conn.Write([]byte("Hello !"))
		}

		request = make([]byte, 128)
	}
}

func main() {
	//解析地址
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:6666")
	if checkErr(err, "ResolveTCPAddr") {
		return
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if checkErr(err, "DialTCP") {
		return
	}

	fmt.Println("Connect server success.")

	gLocker.Lock()
	gCondition = sync.NewCond(&gLocker)

	//发送数据给服务器
	conn.Write([]byte("Hello !"))

	//处理连接（lock在上面调用了，所以clientConnHandler函数必须等wait函数调用后才能lock,这样就能保证调用的先后顺序）
	go clientConnHandler(conn)

	//主线程阻塞，等待Singal结束
	for {
		//条件变量同步等待
		gCondition.Wait()
		break
	}
	gLocker.Unlock()
	fmt.Println("Client finish.")
}
