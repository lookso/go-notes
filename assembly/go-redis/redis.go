package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

func main() {
	fmt.Println("This is a program for go to use go_redis.")
	//connect
	client := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})
	var group sync.WaitGroup
	group.Add(1)
	client.Do("set","myname","toms","ex","5","nx")
	time.Sleep(time.Second*10)
	go func() {
		fmt.Println(123)
		defer group.Done()
		client.Do("set","myname","toms","ex","100","nx")
	}()
	group.Wait()
	
	select {
	
	}


}
