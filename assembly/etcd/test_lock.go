package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	// 客户端配置
	config := clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	}
	// 建立连接
	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	//上锁并创建租约
	lease := clientv3.NewLease(client)

	leaseGrantResp, err := lease.Grant(context.TODO(), 5)
	if err != nil {
		panic(err)
	}
	leaseId := leaseGrantResp.ID
	// 创建一个可取消的租约，主要是为了退出的时候能够释放
	ctx, cancelFunc := context.WithCancel(context.TODO())

	// 释放租约
	defer cancelFunc()
	defer lease.Revoke(context.TODO(), leaseId)

	keepRespChan, err := lease.KeepAlive(ctx, leaseId)
	if err != nil {
		panic(err)
	}
	// 续约应答
	go func() {
		for {
			select {
			case keepResp := <-keepRespChan:
				if keepRespChan == nil {
					fmt.Println("租约已经失效了")
					goto END
				} else { // 每秒会续租一次, 所以就会收到一次应答
					fmt.Println("收到自动续租应答:", keepResp.ID)
				}
			}
		}
	END:
	}()

	// 在租约时间内去抢锁（etcd 里面的锁就是一个 key）
	kv := clientv3.NewKV(client)

	// 创建事务
	txn := kv.Txn(context.TODO())

	// If 不存在 key，Then 设置它，Else 抢锁失败
	txn.If(clientv3.Compare(clientv3.CreateRevision("lock"), "=", 0)).
		Then(clientv3.OpPut("lock", "g", clientv3.WithLease(leaseId))).
		Else(clientv3.OpGet("lock"))

	// 提交事务
	txnResp, err := txn.Commit()
	if err != nil {
		panic(err)
	}

	if !txnResp.Succeeded {
		fmt.Println("锁被占用:", string(txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
		return
	}

	// 抢到锁后执行业务逻辑，没有抢到则退出
	fmt.Println("处理任务")
	time.Sleep(5 * time.Second)
}
