package main

import (
	"context"
	"fmt"
)

type NSQD struct {
	Name     string
	TcpAddr  string
	HttpAddr string
}
type NSQLookupd struct {
	Name     string
	TcpAddr  string
	HttpAddr string
}
type TcpHandle interface {
	Handle() string
}

func TCPServer(ctx context.Context, handle TcpHandle) error {
	return nil
}

type NSQLookupdTcpService struct {
	NSQLookupd NSQLookupd
}

func (p *NSQLookupdTcpService) Handle() string {
	return p.NSQLookupd.Name
}

type NSQdTcpService struct {
	Nsqd NSQD
}

func (p *NSQdTcpService) Handle() string {
	return p.Nsqd.Name
}

func NSQLookupdFunc() {
	ctx := context.Background()
	var err = TCPServer(ctx, &NSQLookupdTcpService{NSQLookupd{
		Name: "nsqloopupd",
	}})
	if err != nil {
		fmt.Println(err)
	}
}

func NsqLookupFun() {
	ctx := context.Background()
	var err = TCPServer(ctx, new(NSQdTcpService))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	NSQLookupdFunc()
	NsqLookupFun()
}
