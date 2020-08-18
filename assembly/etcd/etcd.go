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
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	// 建立客户端成功
	fmt.Println("connect success")
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := cli.Put(ctx, "/config/grpc", "grpc.com")
	if err != nil {
		return err
	}
	return nil
}

func doGet() error {
	if err := mustInit(); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := cli.Get(ctx, "/config", clientv3.WithPrefix())
	if err != nil {
		return err
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
	return nil
}

func doWatch() error {
	var err error
	if err = mustInit(); err != nil {
		return err
	}
	// 监听etcd集群键的改变：
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	wc := cli.Watch(ctx, "/config", clientv3.WithPrefix(), clientv3.WithPrevKV())
	for v := range wc {
		if v.Err() != nil {
			panic(err)
		}
		for _, e := range v.Events {
			fmt.Printf("type:%v\n kv:%v  prevKey:%v  ", e.Type, e.Kv, e.PrevKv)
		}
	}
	return nil
}

func main() {
	conn()
	if err := doSet(); err != nil {
		fmt.Println("doSet", err)
	}
	if err := doGet(); err != nil {
		fmt.Println("doGet", err)
	}
	if err := doWatch(); err != nil {
		fmt.Println("doWatch", err)
	}
	defer cli.Close() // 连接关闭要放在main函数里面
}
