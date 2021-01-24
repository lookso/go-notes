package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmdStr := fmt.Sprintf("ls -l %s | head -n %d", ".", 10)
	cmd := exec.Command("sh", "-c", cmdStr)
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Errorf("Error: %v\n", err)
	} else {
		fmt.Printf("Success: %s\n%s\n", cmdStr, out)
	}

	fmt.Println("---------os/user--------")
	who, _ := user.Current()
	fmt.Printf("Uid:%s,Gid:%s,name:%s\n", who.Uid, who.Gid, who.Name)

	findName, _ := user.Lookup(who.Username)
	fmt.Println(findName.Username)

	fmt.Println("---------os--------")
	fmt.Println(os.Hostname())    // Hostname返回内核提供的主机名。
	fmt.Println(os.Getpagesize()) // Getpagesize返回底层的系统内存页的尺寸。
	fmt.Println(os.Getegid())     // Getuid返回调用者的用户ID。
	fmt.Println(os.Getpid())      // Getpid返回调用者所在进程的进程ID
	fmt.Println(os.Getppid())     // Getppid返回调用者所在进程的父进程的进程ID。

	// Getwd返回一个对应当前工作目录的根路径。如果当前目录可以经过多条路径抵达（因为硬链接），Getwd会返回其中一个
	dir, _ := os.Getwd()
	fmt.Println(dir) // /Users/itech8/data/app/my_app/go/src/go_package/os

	timeFile := dir + "/../time/basic.go"

	file, err := os.Open(timeFile)
	if err != nil {
		errors.New("open file is fail")
	}
	fmt.Println(file.Name())
	text := make([]byte, 100)
	fileRead, _ := file.Read(text)
	fmt.Println(fileRead)
	fmt.Println(string(text)) // 只返回100个字节的内容

	readFile, _ := ioutil.ReadAll(file)
	fmt.Println(string(readFile))

	defer file.Close()

	fmt.Println("---------os/signal--------")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGSTOP, syscall.SIGTERM, syscall.SIGQUIT /*, syscall.SIGKILL*/)
	fmt.Println(<-sigChan)

}
