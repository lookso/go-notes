package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"strconv"
	"sync"
	"time"
)

type EtcdMutex struct {
	Conf    clientv3.Config //etcd集群配置
	lease   clientv3.Lease
	leaseID clientv3.LeaseID
	Ttl     int64              //租约时间
	Key     string             //etcd的key
	cancel  context.CancelFunc //关闭续租的func
	txn     clientv3.Txn
}

func (em *EtcdMutex) init() error {
	var err error
	var ctx context.Context
	client, err := clientv3.New(em.Conf)
	if err != nil {
		return err
	}
	em.txn = clientv3.NewKV(client).Txn(context.TODO())
	em.lease = clientv3.NewLease(client)
	leaseResp, err := em.lease.Grant(context.TODO(), em.Ttl)
	if err != nil {
		return err
	}
	ctx, em.cancel = context.WithCancel(context.TODO())
	em.leaseID = leaseResp.ID
	_, err = em.lease.KeepAlive(ctx, em.leaseID)
	return err
}
func (em *EtcdMutex) Lock() error {
	err := em.init()
	if err != nil {
		return err
	}
	fmt.Println(string(clientv3.OpGet(em.Key).ValueBytes()))
	//LOCK:
	em.txn.If(clientv3.Compare(clientv3.CreateRevision(em.Key), "=", 0)).
		Then(clientv3.OpPut(em.Key, "g", clientv3.WithLease(em.leaseID))).
		Else()
	// 提交事务
	txnResp, err := em.txn.Commit()
	if err != nil {
		return err
	}
	if !txnResp.Succeeded { //判断txn.if条件是否成立
		return fmt.Errorf("锁被占用")
	}
	return nil
}

func (em *EtcdMutex) UnLock(n int) {
	em.cancel()
	em.lease.Revoke(context.TODO(), em.leaseID)
	fmt.Println(strconv.Itoa(n) + "释放了锁")
}

func main() {
	var conf = clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	}
	eMutex1 := &EtcdMutex{
		Conf: conf,
		Ttl:  10,
		Key:  "lock",
	}
	eMutex2 := &EtcdMutex{
		Conf: conf,
		Ttl:  10,
		Key:  "lock",
	}
	var wg sync.WaitGroup
	wg.Add(2)
	//groutine1
	go func(group sync.WaitGroup) {
		err := eMutex1.Lock()
		if err != nil {
			fmt.Println(err)
			fmt.Println("g1抢锁失败")

			return
		}
		fmt.Println("g1抢锁成功")
		time.Sleep(1 * time.Second)
		defer eMutex1.UnLock(1)
		defer group.Done()
	}(wg)

	//groutine2
	go func(group sync.WaitGroup) {
		err := eMutex2.Lock()
		if err != nil {
			fmt.Println(err)
			fmt.Println("g2抢锁失败")
			return
		}
		fmt.Println("g2抢锁成功")
		defer eMutex2.UnLock(2)
		defer group.Done()
	}(wg)
	wg.Wait()
	fmt.Println("----------------------")

}
