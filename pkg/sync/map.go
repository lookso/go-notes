package main

import "sync"

func main() {
	var smp sync.Map
	smp.Load(1)
}
