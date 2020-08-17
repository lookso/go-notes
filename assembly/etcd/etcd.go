/*
@Time : 2019/1/15 5:03 PM
@Author : Tenlu
@File : basic
@Software: GoLand
*/
package main

import (
	"context"
	"errors"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var cli *clientv3.Client

func conn() {
	var err error
	cli, err = clientv3.New(
		clientv3.Config{
			Endpoints:   []string{"localhost:2379"},
			DialTimeout: 5 * time.Second,
		})
	defer cli.Close()
	if err != nil {
		fmt.Println(errors.New("connect failed"))
	}
	fmt.Println("connect Etcd success")
}
func mustInit() error {
	if cli == nil {
		return errors.New("config is not init")
	}
	return nil
}
func doSet() error {
	if err := mustInit(); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	// 设置key="/tizi365/url" 的值为 www.tizi365.com
	_, err := cli.Put(ctx, "/tizi365/url", "www.tizi365.com")
	if err != nil {
		return err
	}
	return nil
}
func main() {
	conn()
	if err := doSet(); err != nil {
		fmt.Println("doGet", err)
	}
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
