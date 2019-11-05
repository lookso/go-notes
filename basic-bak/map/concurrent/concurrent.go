package main

import (
	"strconv"
	"sync"
)

// golang 并发不安全
// M
type M struct {
	Map map[string]string
}

// Set ...
func (m *M) Set(key, value string) {
	m.Map[key] = value
}

// Get ...
func (m *M) Get(key string) string {
	return m.Map[key]
}

func main() {
	c := M{Map: make(map[string]string)}
	var wg sync.WaitGroup
	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func(n int) {
			k, v := strconv.Itoa(n), strconv.Itoa(n)
			c.Set(k, v)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}
