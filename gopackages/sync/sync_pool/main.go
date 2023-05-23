package main

import (
	"fmt"
	"sync"
	_ "unsafe"
)

type Obj struct {
	mu sync.Mutex
}

func (o Obj) Lock()        { o.mu.Lock() }
func (o Obj) Dosomething() {}
func (o Obj) Unlock()      { o.mu.Unlock() }

//go:linkname FastRand runtime.fastrand
//func FastRand() uint32


type noCopy struct{}

// Lock is a no-op used by -copylocks checker from `go vet`.
func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}


//TabNine:config

func main()  {
	xesPool := sync.Pool {
		New: func() interface {} {
			return "new example"
		},
	}
	p:=xesPool
	fmt.Println(p.Get())
	p.Put("abc")
	fmt.Println(p.Get())
	fmt.Println(p.Get())

	fmt.Println(p.Get())
}

