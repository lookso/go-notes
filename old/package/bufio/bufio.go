package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//os.O_RDONLY: 以只读的方式打开
	//os.O_WRONLY: 以只写的方式打开
	//os.O_RDWR : 以读写的方式打开
	//os.O_NONBLOCK: 打开时不阻塞
	//os.O_APPEND: 以追加的方式打开
	//os.O_CREAT: 创建并打开一个新文件
	//os.O_TRUNC: 打开一个文件并截断它的长度为零（必须有写权限）
	//os.O_EXCL: 如果指定的文件存在，返回错误
	//os.O_SHLOCK: 自动获取共享锁
	//os.O_EXLOCK: 自动获取独立锁
	//os.O_DIRECT: 消除或减少缓存效果
	//os.O_FSYNC : 同步写入
	//os.O_NOFOLLOW: 不追踪软链接
	//
	s := strings.NewReader("ACCLIMATISED")
	br := bufio.NewReader(s)
	b := make([]byte, 20)
	n, _ := br.Read(b)
	fmt.Println("写入的字节数:", n)

	// Peek返回输入流的下n个字节，而不会移动读取位置。返回的[]byte只在下一次调用读取操作前合法。
	// 如果Peek返回的切片长度比n小，它也会返会一个错误说明原因。
	// 如果n比缓冲尺寸还大，返回的错误将是ErrBufferFull。
	p,_:=br.Peek(5)
	fmt.Printf("p:%s\n", p)

	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	bufFile := bufio.NewWriter(file)
	num, _ := bufFile.WriteString(time.Now().Format("2006-01-02 15:03:04\n"))
	// 为什么会在文件中看不到写入的数据呢，我们来看看bufio中的Write方法。

	//Write方法首先会判断写入的数据长度是否大于设置的缓冲长度，如果小于，则会将数据copy到缓冲中；
	//当数据长度大于缓冲长度时，如果数据特别大，则会跳过copy环节，直接写入文件。
	//其他情况依然先会将数据拷贝到缓冲队列中，然后再将缓冲中的数据写入到文件中。
	if err = bufFile.Flush(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)
}
