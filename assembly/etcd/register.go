package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"sync"
	"time"
)

// service register and find
type SrvReg struct {
	client        *clientv3.Client
	lease         clientv3.Lease
	leaseResp     *clientv3.LeaseGrantResponse
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	cancleFunc    func()
	wg            *sync.WaitGroup
}

func NewSrvReg(addr []string, ttl int64) (*SrvReg, error) {
	conf := clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 5 * time.Second,
	}
	client, err := clientv3.New(conf)
	if err != nil {
		return nil, err
	}
	srv := &SrvReg{
		client: client,
	}
	srv.SetLease(ttl)
	go srv.ListenLeaseRespChan()
	return srv, nil
}
func (this *SrvReg) SetLease(ttl int64) (err error) {
	lease := clientv3.NewLease(this.client)
	grantResponse, err := lease.Grant(context.TODO(), ttl)
	if err != nil {
		return err
	}
	//设置续租
	ctx, cancelFunc := context.WithCancel(context.TODO())
	leaseRespChan, err := lease.KeepAlive(ctx, grantResponse.ID)

	this.cancleFunc = cancelFunc
	this.lease = lease
	this.leaseResp = grantResponse
	this.keepAliveChan = leaseRespChan
	return err
}

//监听 续租情况
func (this *SrvReg) ListenLeaseRespChan() {
	for {
		select {
		case leaseKeepResp := <-this.keepAliveChan:
			if leaseKeepResp == nil {
				fmt.Printf("已经关闭续租功能\n")
				break
			} else {
				fmt.Printf("续租成功\n")
				return
			}
		default:
			fmt.Println("....")
		}
	}
}

func (this *SrvReg) PutService(key, value string) (err error) {
	kv := clientv3.NewKV(this.client)
	fmt.Println(this.leaseResp.ID)
	_, err = kv.Put(context.TODO(), key, value, clientv3.WithLease(this.leaseResp.ID))
	if err != nil {
		return err
	}
	return nil
}
func main() {
	srv, _ := NewSrvReg([]string{"localhost:12379", "localhost:22379", "localhost:32379"}, 10)
	if err := srv.PutService("/node/111", "hello"); err != nil {
		fmt.Println(err)
	}
	select {}
}
