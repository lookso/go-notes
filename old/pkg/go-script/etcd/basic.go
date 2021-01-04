/*
@Time : 2019/1/15 5:03 PM
@Author : Tenlu
@File : basic
@Software: GoLand
*/
package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	/*
		DialTimeout time.Duration `json:"dial-timeout"`
		Endpoints []string `json:"endpoints"`
	*/
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect etcd success")
	defer cli.Close()
}

// go大数据库日志
// https://studygolang.com/articles/11116
// https://studygolang.com/articles/11396
// https://www.sohu.com/a/168273396_657921
//

// 下载etcd客户端
// go get go.etcd.io/etcd/client
/**
itech8deMacBook-Pro :: /data/db » ETCDCTL_API=3 etcdctl put foo bar                                                                                            127 ↵
OK

itech8deMacBook-Pro :: /data/db » ETCDCTL_API=3 etcdctl get foo
foo
bar

*/
