package util

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

type Options struct {
	waitGroup WaitGroupWrapper
}

func getErr(reason string) error {
	return errors.New("this is err," + reason)
}
func (o *Options) InitOption() {
	exitCh := make(chan error)
	var once sync.Once
	exitFunc := func(err error) {
		once.Do(func() {
			if err != nil {
				fmt.Errorf("%s", err)
			}
			fmt.Println("err:", err)
			exitCh <- err
		})
	}
	o.waitGroup.Wrap(func() {
		exitFunc(getErr("init-456"))
	})
	o.waitGroup.Wrap(func() {
		exitFunc(getErr("init-123"))
	})
	fmt.Println(<-exitCh)
	o.waitGroup.Wait()
}

func TestWaitGroupWrapper(t *testing.T) {
	new(Options).InitOption()
}
