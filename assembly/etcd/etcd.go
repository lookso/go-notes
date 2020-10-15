package main

import (
	"context"
	"errors"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

var cli *clientv3.Client

var (
	LeaseKeepAliveFail = errors.New("lease keepAlive fail")
)

func conn() {
	var err error
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:112379", "localhost:122379", "localhost:132379"},
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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	//withPrevKV()是为了获取操作前已经有的key-value
	putResp, err := cli.Put(ctx, "/config/golang", "golang.com", clientv3.WithPrevKV())
	if err != nil {
		return err
	}
	if putResp.PrevKv != nil {
		fmt.Println("old:", string(putResp.PrevKv.Key), string(putResp.PrevKv.Value))
		_, err := cli.Put(ctx, "/config/etcd", "etcd.io")
		if err != nil {
			return err
		}
	}
	return nil
}

func doGet() error {
	if err := mustInit(); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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
//监听某个key的变化
func doWatch() error {
	group := new(errgroup.Group)
	var err error
	if err = mustInit(); err != nil {
		return err
	}
	// 监听etcd集群键的改变：
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	group.Go(func() error {
		wc := cli.Watch(ctx, "/config", clientv3.WithPrefix(), clientv3.WithPrevKV())
		for v := range wc {
			if v.Err() != nil {
				return errors.New("watch key err")
			}
			for _, e := range v.Events {
				fmt.Printf("type:%v\n kv:%v  prevKey:%v  ", e.Type, e.Kv, e.PrevKv)
			}
		}
		return err
	})
	if err := group.Wait(); err != nil {
		return err
	}
	return nil
}
func doDel() error {
	var err error
	if err = mustInit(); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = cli.Delete(ctx, "/config/etcd",clientv3.WithPrefix())
	if err != nil {
		return err
	}
	return nil
}

func doLease() error {
	var err error
	if err = mustInit(); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 创建租约
	lease := clientv3.NewLease(cli)
	leaseResp, err := lease.Grant(ctx, 10)
	if err != nil {
		return err
	}
	// 指定租约
	kv := clientv3.NewKV(cli)
	_, err = kv.Put(context.TODO(), "golang", "golang.com", clientv3.WithLease(leaseResp.ID))
	if err != nil {
		return err
	}
	//续租
	leaseRespChan, err := lease.KeepAlive(context.TODO(), leaseResp.ID)
	if err != nil {
		return LeaseKeepAliveFail
	}
	var group sync.WaitGroup
	group.Add(1)

	// 监听租约
	go func(leaseRespChan <-chan *clientv3.LeaseKeepAliveResponse) {
		defer group.Done()
		for {
			select {
			case leaseKeepResp := <-leaseRespChan:
				if leaseKeepResp == nil {
					fmt.Printf("keepalive lease fail\n")
					break
				} else {
					fmt.Printf("keepalive lease success\n")
					return
				}
			}
		}
	}(leaseRespChan)
	group.Wait()
	//defer逻辑可以保证租约被清理，防止长期占用key
	defer lease.Revoke(ctx,leaseResp.ID)
	return nil
}

func main() {
	conn()
	for {
		if err := doSet(); err != nil {
			fmt.Println("doSet", err)
		}
		if err := doGet(); err != nil {
			fmt.Println("doGet", err)
		}
		if err := doDel(); err != nil {
			fmt.Println("doDel", err)
		}
		if err := doLease(); err != nil {
			fmt.Println("doLease", err)
		}
		if err := doWatch(); err != nil {
			fmt.Println("doWatch", err)
		}
		fmt.Println("---------------")
		
		time.Sleep(1 * time.Second)
	}
	defer cli.Close() // 连接关闭要放在main函数里面
}
